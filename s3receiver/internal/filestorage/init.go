package filestorage

import (
	"context"
	"io"
	"sync"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v6"
)

const _N = 6

type Filestorage struct {
	conf *FilestorageConfig
	cli  *minio.Client
}

func New(conf *FilestorageConfig) (*Filestorage, error) {
	cli, err := minio.New(conf.Address, conf.Creds.Username, conf.Creds.Password, false)
	if err != nil {
		return nil, err
	}
	return &Filestorage{
		conf: conf,
		cli:  cli,
	}, nil
}

func (fs *Filestorage) GetOne(ctx context.Context, bucket uuid.UUID, id uuid.UUID) (<-chan []byte, error) {
	ctx, cancel := context.WithTimeout(ctx, fs.conf.Timeout)
	ok, err := fs.cli.BucketExists(bucket.String())
	if !ok {
		cancel()
		return nil, ErrBucketNotExist
	}
	if err != nil {
		cancel()
		return nil, err
	}
	obj, err := fs.cli.GetObject(bucket.String(), id.String(), minio.GetObjectOptions{})
	if err != nil {
		cancel()
		return nil, err
	}
	var respChan = make(chan []byte, _N)

	go func(file *minio.Object, ch chan<- []byte) {
		defer file.Close()
		defer close(ch)
		defer cancel()

		buf := make([]byte, fs.conf.ChannelPayloadSize)

		for {
			n, err := file.Read(buf)
			if n > 0 {
				chunk := make([]byte, n)
				copy(chunk, buf[:n])

				select {
				case ch <- chunk:
				case <-ctx.Done():
					return
				}
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				break
			}
		}
	}(obj, respChan)

	return respChan, nil
}

func (fs *Filestorage) GetMany(ctx context.Context, bucket uuid.UUID, ids []uuid.UUID) (<-chan <-chan []byte, error) {
	ok, err := fs.cli.BucketExists(bucket.String())
	if !ok {
		return nil, ErrBucketNotExist
	}
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(ids))

	respChanArr := make(chan (<-chan []byte), len(ids))

	go func() {
		wg.Wait()
		defer close(respChanArr)
	}()

	for _, id := range ids {
		go func(id uuid.UUID) {

			ctx, cancel := context.WithTimeout(ctx, fs.conf.Timeout)
			obj, err := fs.cli.GetObject(bucket.String(), id.String(), minio.GetObjectOptions{})
			if err != nil {
				cancel()
				return
			}
			var respChan = make(chan []byte, _N)
			respChanArr <- respChan

			go func(file *minio.Object, ch chan<- []byte) {
				defer file.Close()
				defer close(ch)
				defer cancel()

				buf := make([]byte, fs.conf.ChannelPayloadSize)

				for {
					n, err := file.Read(buf)
					if n > 0 {
						chunk := make([]byte, n)
						copy(chunk, buf[:n])

						select {
						case ch <- chunk:
						case <-ctx.Done():
							return
						}
					}
					if err != nil {
						if err == io.EOF {
							break
						}
						break
					}
				}
			}(obj, respChan)
		}(id)
	}
	return respChanArr, nil
}
