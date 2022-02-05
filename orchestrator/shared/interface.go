package shared
import (
	"github.com/hashicorp/go-plugin"
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/orchestrator"
	// "net/rpc"
)


type Scrapper interface {
	ScrapAll(Title string) (model.Comic, error)
	ScrapPerChapter(Title, Id string) (model.Chapter, error)
}

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "scrap",
}


// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"scrapper": &ScrapperPlugin{},
}


type ScrapperPlugin struct {
	Impl Scrapper
}
func (p *ScrapperPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (*ScrapperPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

