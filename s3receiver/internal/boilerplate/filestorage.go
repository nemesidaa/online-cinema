package boilerplate

import (
	"github.com/google/uuid"
)

type FilestorageImplementation interface {
	GetOne(bucket uuid.UUID, id uuid.UUID) (chan byte, error)
	GetMany(bucket uuid.UUID, ids []uuid.UUID) ([]chan byte, error)
}
