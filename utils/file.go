package utils

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
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

func GetRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
