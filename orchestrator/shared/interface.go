package shared
import (
	"github.com/ananrafs1/gomic/model"
)


type Scrapper interface {
	ScrapAll(Title string) model.Comic
	ScrapPerChapter(Title, Id string) model.Chapter
}
