package orchestrator

import (
	"log"
	"os/exec"
	"path/filepath"

	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/utils"
	"github.com/hashicorp/go-plugin"
)

type Orchestrator struct {
	PluginName string
	PluginType string
}

func (o *Orchestrator) Orchest(pluginName string) (interface{}, func(), error) {
	fileName := filepath.Join(utils.GetRootDir(), model.Config.Plugins[o.PluginType], pluginName)
	if exist, err := utils.IsExists(fileName); !exist {
		return nil, nil, err
	}
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: Handshake,
		Plugins:         PluginMap,
		Cmd:             exec.Command(fileName),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC},
	})

	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	raw, err := rpcClient.Dispense(o.PluginName)
	if err != nil {
		log.Fatal(err)
	}
	return raw, func() { client.Kill() }, nil
}
