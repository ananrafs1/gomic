package scrapper

import (
	"github.com/ananrafs1/gomic/orchestrator"
	"github.com/ananrafs1/gomic/model"
	"fmt"
)

func ScrapAll(Host, Title string) (model.Comic, error) {
	Hst, close, err := orchestrator.Orchest(Host)
	if err != nil {
		return model.Comic{}, err
	}
	defer close()
	return Hst.ScrapAll(Title)
}
