package scrapper

import (
	"fmt"
	"log"
	"strings"

	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/orchestrator"
	plugin "github.com/ananrafs1/gomic/orchestrator/shared/scrapper"
)

func ScrapAll(Host, Title string, Page, Quantity int) *model.Comic {
	validate(&Host)
	orchestrator := orchestrator.Orchestrator{PluginName: "grpcscrapper", PluginType: "scrapper"}
	Hst, close, err := orchestrator.Orchest(Host)
	if err != nil {
		log.Fatal(err, "not found")
		return nil
	}
	defer close()
	ret := Hst.(plugin.Scrapper).Scrap(Title, Page, Quantity)
	return &ret
}
func validate(host *string) {
	ext := "exe"
	splitByDot := strings.Split(*host, ".")
	if len(splitByDot) > 0 && splitByDot[len(splitByDot)-1] == ext {
		return
	}
	*host += fmt.Sprintf(".%s", ext)
}
