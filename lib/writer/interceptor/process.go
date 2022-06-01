package interceptor

import (
	"context"

	"github.com/ananrafs1/gomic/lib/writer"
	"github.com/ananrafs1/gomic/model"
)

type WriterProcess struct {
	WR writer.IWriter
	p  model.Process
}

func (wp *WriterProcess) Store(ctx context.Context, Image model.Image, ComicInfo model.ComicInfo) error {
	if wp.p.Finish != nil {
		defer wp.p.Finish()
	}
	if wp.p.Start != nil {
		wp.p.Start()
	}
	return wp.WR.Store(ctx, Image, ComicInfo)
}

func (wp *WriterProcess) OnStart(onStartFunc func()) {
	wp.p.Start = onStartFunc
}

func (wp *WriterProcess) OnFinished(onFinishFunc func()) {
	wp.p.Finish = onFinishFunc
}
