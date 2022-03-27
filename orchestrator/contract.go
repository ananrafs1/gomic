package orchestrator

import (
	pg "github.com/ananrafs1/gomic/orchestrator/shared/scrapper"
	"github.com/hashicorp/go-plugin"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "scrap",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	// "scrapper": &pg.ScrapperPlugin{},
	"grpcscrapper": &pg.GRPCPlugin{},
}
