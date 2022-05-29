package main

import (
	"context"
	"flag"
	"log"

	"github.com/ananrafs1/gomic/core"
	"github.com/ananrafs1/gomic/lib/writer/downloader"
	"github.com/ananrafs1/gomic/model"
)

var (
	host, title    string
	Page, Quantity int
)

func main() {
	err := core.Process(context.Background(), host, title, Page, Quantity, &downloader.Downloader{})
	log.Fatal(err)
}

func init() {
	model.InitAndListenConfig()
	readParams()
}
func readParams() {
	flag.StringVar(&host, "host", "template", "host")
	flag.StringVar(&title, "title", "random", "title/chapter")
	flag.IntVar(&Page, "Page", 1, "Page")
	flag.IntVar(&Quantity, "Quantity", 10, "Quantity per Page")
	flag.Parse()
}
