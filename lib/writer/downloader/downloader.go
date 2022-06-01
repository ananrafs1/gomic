package downloader

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/utils"
)

type Downloader struct {
}

func (d *Downloader) Store(ctx context.Context, Image model.Image, ComicInfo model.ComicInfo) error {
	dir := filepath.Join(model.Config.OutputDir, ComicInfo.Title.Name, ComicInfo.Chapter.Id)
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
