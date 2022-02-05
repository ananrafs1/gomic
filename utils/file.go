package utils

import (
	"os"
	"errors"
	"io/ioutil"
	"log"
)

func IsExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func Print(filepath string) error {
	
	if exist, err := IsExists(filepath); !exist {
		return err
	}
	b, err := ioutil.ReadFile(filepath) 
    if err != nil {
        log.Println(err)
    }
    str := string(b) 
    log.Println(str) 
	return nil
}
