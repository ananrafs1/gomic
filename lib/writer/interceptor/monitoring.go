package interceptor

import (
	"context"

	"github.com/ananrafs1/gomic/lib/monitoring"
	"github.com/ananrafs1/gomic/lib/writer"
	"github.com/ananrafs1/gomic/model"
	"go.opentelemetry.io/otel/label"
)

type WriterMonitor struct {
	WR writer.IWriter
}

func (wm *WriterMonitor) Store(ctx context.Context, Image model.Image, ComicInfo model.ComicInfo) error {
	span, closer, ctx := monitoring.CreateSpan(ctx, "Store")
	defer closer()
	span.SetAttributes(label.String("Title", ComicInfo.Title.Name), label.String("Chapter", ComicInfo.Chapter.Id), label.Int("Episode", Image.Episode))
	return wm.WR.Store(ctx, Image, ComicInfo)
}
