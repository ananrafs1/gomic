package interceptor

import (
	"context"

	"github.com/ananrafs1/gomic/lib/writer"
	"github.com/ananrafs1/gomic/model"
)

type WriterProcess struct {
	WR writer.IWriter
	model.Process
}

func (wp *WriterProcess) Store(ctx context.Context, Image model.Image, ComicInfo model.ComicInfo) error {
	if wp.Finish != nil {
		defer wp.Finish()
	}
	if wp.Start != nil {
		wp.Start()
	}
	return wp.WR.Store(ctx, Image, ComicInfo)
}

func (wp *WriterProcess) OnStart(onStartFunc func()) {
	wp.Start = onStartFunc
}

func (wp *WriterProcess) OnFinished(onFinishFunc func()) {
	wp.Finish = onFinishFunc
}
