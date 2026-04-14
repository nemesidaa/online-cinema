package filestorage

import (
	"github.com/google/uuid"
	"github.com/minio/minio-go/v6"
)

type Filestorage struct {
	conf *FilestorageConfig
	cli  *minio.Client
}

// ! Reference:
// type FilestorageImplementation interface {
// 	GetOne(uuid.UUID) (chan byte, error)
// 	GetMany([]uuid.UUID) ([]chan byte, error)
// }

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

func (fs *Filestorage) GetOne(bucket uuid.UUID, id uuid.UUID) (chan byte, error)
func (fs *Filestorage) GetMany(bucket uuid.UUID, ids []uuid.UUID) ([]chan byte, error)
