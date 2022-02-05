package orchestrator
import (
	"github.com/hashicorp/go-plugin"
	"github.com/ananrafs1/gomic/orchestrator/shared"
	"github.com/ananrafs1/gomic/utils"
	"os/exec"
	"log"
	"path/filepath"
)

func Orchest(Host string) (shared.Scrapper, func(),error) {
	nameFile := filepath.Join(".","plugins",Host)
	if exist, err := utils.IsExists(nameFile); !exist {
		return nil, nil,  err
	}
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command(nameFile),
		
	})
	
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}
	
	raw, err := rpcClient.Dispense("scrapper")
	if err != nil {
		log.Fatal(err)
	}
	return raw.(shared.Scrapper), func() {client.Kill()}, nil
}