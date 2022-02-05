package main
import (
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/scrapper"
	"github.com/ananrafs1/gomic/utils"
	"log"
	"fmt"
	"path/filepath"
)

func main() {


	
}

 func init(){
	 go ListenConfigurationChange()

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