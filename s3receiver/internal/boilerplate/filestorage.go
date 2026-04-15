package boilerplate

import (
	"context"

	"github.com/google/uuid"
)

type FilestorageImplementation interface {
	GetOne(ctx context.Context, bucket uuid.UUID, id uuid.UUID) (<-chan []byte, error)
	GetMany(ctx context.Context, bucket uuid.UUID, ids []uuid.UUID) (<-chan <-chan []byte, error)
}
