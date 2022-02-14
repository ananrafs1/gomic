package scrapper

import (
	"github.com/ananrafs1/gomic/orchestrator"
	"github.com/ananrafs1/gomic/model"
	"strings"
	"fmt"
	"log"
)

func ScrapAll(Host, Title string) *model.Comic {
	validate(&Host)
	Hst, close, err := orchestrator.Orchest(Host)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer close()
	ret := Hst.Scrap(Title, 1, 10)
	return &ret
}
func validate(host *string)  {
	ext := "exe"
	splitByDot := strings.Split(*host, ".")
	if len(splitByDot) > 0 && splitByDot[len(splitByDot)-1] == ext {
		return
	}
	*host += fmt.Sprintf(".%s", ext)
}