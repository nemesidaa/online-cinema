package boilerplate

import (
	"context"
	"tpserver/internal/model"
)

type ControllerImplementation interface {
	GetThumb(context.Context, *model.GetThumbRequest) ([]byte, error)
}
