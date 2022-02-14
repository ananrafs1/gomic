package orchestrator
import (
	"github.com/hashicorp/go-plugin"
	"github.com/ananrafs1/gomic/utils"
	pg "github.com/ananrafs1/gomic/orchestrator/shared/plugin"
	"os/exec"
	"log"
	"path/filepath"
)

func Orchest(Host string) (pg.Scrapper, func(),error) {
	nameFile := filepath.Join(".","plugins",Host)
	if exist, err := utils.IsExists(nameFile); !exist {
		return nil, nil,  err
	}
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: Handshake,
		Plugins:         PluginMap,
		Cmd:             exec.Command(nameFile),
		AllowedProtocols: []plugin.Protocol{
			 plugin.ProtocolGRPC},
		
	})
	
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}
	
	raw, err := rpcClient.Dispense("grpcscrapper")
	if err != nil {
		log.Fatal(err)
	}
	return raw.(pg.Scrapper), func() {client.Kill()}, nil
}