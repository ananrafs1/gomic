package writer

import(
	"github.com/ananrafs1/gomic/model"
)

type IWriter interface {
	Store(Image model.Image, ComicInfo model.ComicInfo) error
	OnStart(onStart func())
	OnFinished(onStart func())
}