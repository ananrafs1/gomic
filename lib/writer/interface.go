package writer

import (
	"context"

	"github.com/ananrafs1/gomic/model"
)

type IWriter interface {
	Store(ctx context.Context, Image model.Image, ComicInfo model.ComicInfo) error
	OnStart(onStart func())
	OnFinished(onStart func())
}
