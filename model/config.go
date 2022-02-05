package model
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/fsnotify/fsnotify"
)
type Configuration struct {
	OutputDir string
}

func (c *Configuration) ParseConfig(filepath string) error {
	raw, err := ioutil.ReadFile(filepath)
    if err != nil {
       log.Println("Error occured while reading config")
       return err
    }
    json.Unmarshal(raw, c)
	return nil
}

func WatchConfig(filepath string, changes chan string, errorNotif chan error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	defer watcher.Close()
	err = watcher.Add(filepath)
	if err !=  nil {
		return
	}
	done := make(chan bool)
	go func(){
		changes <- filepath
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				
				if event.Op&fsnotify.Write == fsnotify.Write {
					changes <- event.Name
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
				errorNotif <- err
			}
		}
	}()
	<- done
}

var Config Configuration