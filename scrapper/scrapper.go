package scrapper

import (
	"github.com/ananrafs1/gomic/orchestrator"
	"github.com/ananrafs1/gomic/model"
	"strings"
	"fmt"
)

func ScrapAll(Host, Title string) (model.Comic, error) {
	validate(&Host)
	Hst, close, err := orchestrator.Orchest(Host)
	if err != nil {
		return model.Comic{}, err
	}
	defer close()
	return Hst.ScrapAll(Title)
}
func validate(host *string)  {
	ext := "exe"
	splitByDot := strings.Split(*host, ".")
	if len(splitByDot) > 0 && splitByDot[len(splitByDot)-1] == ext {
		return
	}
	*host += fmt.Sprintf(".%s", ext)
}