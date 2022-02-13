package main
import (
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/utils"
	"github.com/ananrafs1/gomic/core"
	"github.com/ananrafs1/gomic/writer/downloader"
	"log"
	"path/filepath"
	"flag"
)
var (
	host, title string
)

func main() {
	err := core.Proces(host, title, &downloader.Downloader{})
	log.Fatal(err)
}

 func init(){
	 go ListenConfigurationChange()
	 readParams()

 }

 func ListenConfigurationChange() {
	 nameFile := filepath.Join(".","config","Config.json")
	 if exist, _ := utils.IsExists(nameFile); !exist {
		 log.Fatal("File Not Found")
		 return
	 }
	 changes, errNotif := make(chan string), make(chan error)
	 go model.WatchConfig(nameFile, changes, errNotif)
	
	 for {
		 select {
		 case file := <- changes :
			err := model.Config.ParseConfig(file)
			if err != nil {
				log.Fatal(err)
			}

		case errs := <- errNotif :
			log.Fatal("Error when Reading Configuration Changes, err: ", errs)
		 }
	 }
 }

 func readParams(){
	flag.StringVar(&host, "host", "template", "host")
	flag.StringVar(&title, "title", "random", "title/chapter")
	flag.Parse()
 }