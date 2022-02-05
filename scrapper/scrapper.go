package scrapper

import (
	"github.com/hashicorp/go-plugin"
	"github.com/ananrafs1/gomic/orchestrator/shared"
	"github.com/ananrafs1/gomic/utils"
	"github.com/ananrafs1/gomic/model"
	"os/exec"
	"log"
	"path/filepath"
)

func ScrapAll(Host, Title string) (model.Comic, error) {
	nameFile := filepath.Join(".","plugins",Host)
	if exist, err := utils.IsExists(nameFile); !exist {
		
		return model.Comic{}, err
	}
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command(nameFile),
		
	})
	defer client.Kill()
	
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}
	
	raw, err := rpcClient.Dispense("scrapper")
	if err != nil {
		log.Fatal(err)
	}
	scrap := raw.(shared.Scrapper)
	return scrap.ScrapAll(Title)
}
