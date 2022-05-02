package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ananrafs1/gomic/utils"
	"github.com/fsnotify/fsnotify"
)

type Configuration struct {
	OutputDir string            `json:"outputdir"`
	Plugins   map[string]string `json:"plugins"`
	Server    map[string]string `json:"server"`
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
	if err != nil {
		return
	}
	done := make(chan bool)
	go func() {
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
	<-done
}

func InitAndListenConfig() {
	fileName := filepath.Join(".", "config", "Config.json")
	if exist, _ := utils.IsExists(fileName); !exist {
		log.Fatal("File Not Found")
		return
	}
	Config.ParseConfig(fileName)
	go ListenConfigurationChange(fileName)
}

func ListenConfigurationChange(fileName string) {
	changes, errNotif := make(chan string), make(chan error)
	go WatchConfig(fileName, changes, errNotif)

	for {
		select {
		case file := <-changes:
			err := Config.ParseConfig(file)
			if err != nil {
				log.Fatal(err)
			}

		case errs := <-errNotif:
			log.Fatal("Error when Reading Configuration Changes, err: ", errs)
		}
	}
}

var Config Configuration
