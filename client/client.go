package client

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/sha256"
	"errors"
	"github.com/klauspost/compress/zstd"
	"github.com/tigerwill90/webcb/internal/crypto"
	"github.com/tigerwill90/webcb/internal/enclave"
	grpchelper "github.com/tigerwill90/webcb/internal/grpc"
	"github.com/tigerwill90/webcb/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"math/rand"
	"time"
)

type Client struct {
	c           proto.WebClipboardClient
	chunkSize   int64
	checksum    bool
	ttl         time.Duration
	compression bool
	password    []byte
}

func New(conn grpc.ClientConnInterface, opts ...Option) *Client {
	config := defaultConfig()
	for _, opt := range opts {
		opt(config)
	}

	return &Client{
		c:           proto.NewWebClipboardClient(conn),
		chunkSize:   config.chunkSize,
		checksum:    config.checksum,
		ttl:         config.ttl,
		compression: config.compression,
		password:    config.password,
	}
}

func (c *Client) Copy(ctx context.Context, r io.Reader) error {
	stream, err := c.c.Copy(ctx)
	if err != nil {
		return err
	}

	hasher := sha256.New()
	defer hasher.Reset()

	gw := grpchelper.NewWriter(newGrpcSender(stream), make([]byte, c.chunkSize))
	w := io.Writer(gw)
	var enc *zstd.Encoder
	if c.compression {
		enc, err = zstd.NewWriter(w, zstd.WithEncoderLevel(zstd.SpeedFastest))
		if err != nil {
			return err
		}
		w = enc
		defer enc.Close()
	}

	var salt []byte
	var iv []byte
	if len(c.password) > 0 {
		randSalt := make([]byte, crypto.SaltSize)
		rand.Read(randSalt)
		randIv := make([]byte, aes.BlockSize)
		rand.Read(randIv)

		salt = randSalt
		iv = randIv

		e := enclave.New(c.password)
		pwd, destroy := e.Open()
		defer destroy()

		key, err := crypto.DeriveKey(pwd, salt)
		if err != nil {
			return err
		}
		destroy()

		sw, err := crypto.NewStreamWriter(key, iv, w)
		if err != nil {
			return err
		}
		defer sw.Close()
		w = sw
	}

	if err := sendInfo(stream, c.ttl, c.compression, iv, salt); err != nil {
		return err
	}

	buf := make([]byte, c.chunkSize)
	for {
		n, err := io.ReadAtLeast(r, buf, int(c.chunkSize))
		if err != nil {
			// The error is EOF only if no bytes were read
			if errors.Is(err, io.EOF) {
				break
			}
			// If an EOF happens after reading fewer than min bytes,
			// ReadAtLeast returns ErrUnexpectedEOF
			if errors.Is(err, io.ErrUnexpectedEOF) {
				if _, err := w.Write(buf[:n]); err != nil {
					return err
				}
				if enc != nil {
					if err := enc.Close(); err != nil {
						return err
					}
				}
				if c.checksum {
					if _, err := hasher.Write(buf[:n]); err != nil {
						return err
					}
				}
				continue
			}

			return err
		}

		if _, err := w.Write(buf[:n]); err != nil {
			return err
		}
		if c.checksum {
			if _, err := hasher.Write(buf[:n]); err != nil {
				return err
			}
		}
	}

	var checksum []byte
	if c.checksum {
		checksum = hasher.Sum(nil)
	}
	if err := gw.Flush(checksum); err != nil {
		return err
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Paste(ctx context.Context, w io.Writer) error {
	resp, err := c.c.Paste(ctx, &proto.PasteOption{Length: c.chunkSize}, grpc.MaxCallRecvMsgSize(int(c.chunkSize+1000)))
	if err != nil {
		return err
	}

	hasher := sha256.New()
	defer hasher.Reset()

	stream, err := resp.Recv()
	if err != nil {
		return err
	}

	info := stream.GetInfo()
	if info == nil {
		return errors.New("protocol error: chunk stream expected but get info header")
	}

	if len(c.password) == 0 && (len(info.Iv) != 0 || len(info.Salt) != 0) {
		return errors.New("data is encrypted but no password provided")
	}

	gr := grpchelper.NewReader(newGrpcReceiver(resp))
	r := io.Reader(gr)
	if info.Compressed {
		dec, err := zstd.NewReader(r)
		if err != nil {
			return err
		}
		defer dec.Close()
		r = dec
	}
	if len(c.password) > 0 && len(info.Iv) == aes.BlockSize && len(info.Salt) == crypto.SaltSize {
		e := enclave.New(c.password)
		pwd, destroy := e.Open()
		defer destroy()

		key, err := crypto.DeriveKey(pwd, info.Salt)
		if err != nil {
			return err
		}
		destroy()

		sr, err := crypto.NewStreamReader(key, info.Iv, r)
		if err != nil {
			return err
		}
		r = sr
	}

	buf := make([]byte, c.chunkSize)
	for {
		n, err := r.Read(buf)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return err
			}

			if n > 0 {
				if info.Checksum {
					if _, err := hasher.Write(buf[:n]); err != nil {
						return err
					}
				}
				if _, err := w.Write(buf[:n]); err != nil {
					return err
				}
			}

			break
		}

		if info.Checksum {
			if _, err := hasher.Write(buf[:n]); err != nil {
				return err
			}
		}
		if _, err := w.Write(buf[:n]); err != nil {
			return err
		}

	}

	if info.Checksum {
		if len(gr.Checksum()) != sha256.Size {
			return errors.New("invalid sha256 size")
		}
		if n := bytes.Compare(gr.Checksum(), hasher.Sum(nil)); n != 0 {
			return errors.New("checksum failed")
		}
	}

	return nil
}

func (c *Client) Clean(ctx context.Context) error {
	_, err := c.c.Clean(ctx, &emptypb.Empty{})
	return err
}

func sendInfo(stream proto.WebClipboard_CopyClient, ttl time.Duration, compressed bool, iv []byte, salt []byte) error {
	return stream.Send(&proto.CopyStream{Data: &proto.CopyStream_Info_{
		Info: &proto.CopyStream_Info{
			Ttl:        int64(ttl),
			Compressed: compressed,
			Salt:       salt,
			Iv:         iv,
		},
	}})
}

type grpcSender struct {
	c proto.WebClipboard_CopyClient
}

func newGrpcSender(client proto.WebClipboard_CopyClient) *grpcSender {
	return &grpcSender{client}
}

func (gc *grpcSender) SendChunk(p []byte) error {
	return gc.c.Send(&proto.CopyStream{Data: &proto.CopyStream_Chunk{
		Chunk: p,
	}})
}

func (gc *grpcSender) SendChecksum(p []byte) error {
	if len(p) == 0 {
		return nil
	}
	return gc.c.Send(&proto.CopyStream{Data: &proto.CopyStream_Checksum{
		Checksum: p,
	}})
}

func newGrpcReceiver(client proto.WebClipboard_PasteClient) *grpcReceiver {
	return &grpcReceiver{c: client}
}

type grpcReceiver struct {
	c        proto.WebClipboard_PasteClient
	checksum []byte
}

func (gc *grpcReceiver) Next() ([]byte, error) {
	stream, err := gc.c.Recv()
	if err != nil {
		return nil, err
	}

	switch stream.Data.(type) {
	case *proto.PastStream_Info_:
		return nil, errors.New("protocol error: chunk stream expected but get info header")
	case *proto.PastStream_Checksum:
		gc.checksum = stream.GetChecksum()
		return nil, io.EOF
	default:
	}
	return stream.GetChunk(), nil
}

func (gc *grpcReceiver) Checksum() []byte {
	return gc.checksum
}
