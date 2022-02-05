package writer

import(
	"github.com/ananrafs1/gomic/model"
)

type Writer interface {
	Store(Image model.Image, ComicInfo model.ComicInfo) error
}