package shared
import (
	// "github.com/hashicorp/go-plugin"
	"github.com/ananrafs1/gomic/model"
	// "net/rpc"
)


type Scrapper interface {
	ScrapAll(Title string) (model.Comic, error)
	ScrapPerChapter(Title, Id string) (model.Chapter, error)
}
