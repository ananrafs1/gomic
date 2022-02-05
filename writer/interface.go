package writer

import(
	"gitlab.com/appdev.ananrafs1/go-micrenderer/model"
)

type Writer interface {
	Store(Image model.Image, ComicInfo model.ComicInfo) error
}