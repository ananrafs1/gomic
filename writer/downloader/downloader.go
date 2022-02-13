package downloader

import (
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/utils"
	"os"
	"path/filepath"
	"strconv"
	"fmt"
	"net/http"
	"io"
)

type Downloader struct{
	model.Process
}

func (d *Downloader) Store(Image model.Image, ComicInfo model.ComicInfo) error {
	defer func(){
		if d.Finish != nil {
			d.Finish()
		}
	}()
	if d.Start != nil {
		d.Start()
	}
	dir := filepath.Join(model.Config.OutputDir, ComicInfo.Title.Name)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	nameFile := filepath.Join(dir, strconv.Itoa(Image.Episode))
	if exist, err := utils.IsExists(fmt.Sprintf(`%s.jpg`, nameFile)); exist {
		// return errors.New("file exist")
		return err
	}
	var response *http.Response
	for _, v := range Image.Link {
		response, err = http.Get(v)
		if err != nil {
			continue
		}
		defer response.Body.Close()
		if response.StatusCode == http.StatusOK {
			break
		}
	}
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] [%s]", response.StatusCode, Image.Link)
		return err
	}

	file, err := os.Create(fmt.Sprintf(`%s.jpg`, nameFile))
	if err != nil {
		err = fmt.Errorf("[os.Create] [%s]", err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		err = fmt.Errorf("[io.Copy] [%s]", err)
		return err
	}

	return nil


}

func (d *Downloader) OnStart(onStart func()) {
	(*d).Start = onStart
}

func (d *Downloader) OnFinished(onFinished func()) {
	(*d).Finish = onFinished
}