package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/ananrafs1/gomic/core"
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/utils"
	"github.com/ananrafs1/gomic/writer/downloader"
)

var (
	host, title    string
	Page, Quantity int
)

func main() {
	err := core.Process(host, title, Page, Quantity, &downloader.Downloader{})
	log.Fatal(err)
}

func init() {
	fileName := filepath.Join(".", "config", "Config.json")
	if exist, _ := utils.IsExists(fileName); !exist {
		log.Fatal("File Not Found")
		return
	}
	model.Config.ParseConfig(fileName)
	go ListenConfigurationChange(fileName)
	readParams()

}

func ListenConfigurationChange(fileName string) {
	changes, errNotif := make(chan string), make(chan error)
	go model.WatchConfig(fileName, changes, errNotif)

	for {
		select {
		case file := <-changes:
			err := model.Config.ParseConfig(file)
			if err != nil {
				log.Fatal(err)
			}

		case errs := <-errNotif:
			log.Fatal("Error when Reading Configuration Changes, err: ", errs)
		}
	}
}

func readParams() {
	flag.StringVar(&host, "host", "template", "host")
	flag.StringVar(&title, "title", "random", "title/chapter")
	flag.IntVar(&Page, "Page", 1, "Page")
	flag.IntVar(&Quantity, "Quantity", 10, "Quantity per Page")
	flag.Parse()
}
