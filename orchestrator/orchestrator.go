package orchestrator
import (
	"github.com/hashicorp/go-plugin"
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/orchestrator/shared"
	"net/rpc"
)



type RPCClient struct{ client *rpc.Client }

func (m *RPCClient) ScrapAll(Title string) (model.Comic, error) {
	// We don't expect a response, so we can just use interface{}
	var resp model.Comic
	err := m.client.Call("Plugin.ScrapAll", Title, &resp)

	// The args are just going to be a map. A struct could be better.
	return resp, err
}

func (m *RPCClient) ScrapPerChapter(Title, Id string) (model.Chapter, error) {
	var resp model.Chapter
	err := m.client.Call("Plugin.ScrapPerChapter", map[string]interface{}{
		"Title" : Title, "Id" : Id,
	}, &resp)
	return resp, err
}

type RPCServer struct {
	// This is the real implementation
	Impl shared.Scrapper
}

func (m *RPCServer) ScrapPerChapter(args map[string]interface{}, resp *model.Chapter) error {
	v, err :=  m.Impl.ScrapPerChapter(args["Title"].(string), args["Id"].(string))
	*resp = v
	return err
}

func (m *RPCServer) ScrapAll(Title string, resp *model.Comic) error {
	v, err := m.Impl.ScrapAll(Title)
	*resp = v
	return err
}