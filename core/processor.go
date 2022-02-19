package core

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ananrafs1/go-palugada/retry"
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/scrapper"
	"github.com/ananrafs1/gomic/writer"
	"github.com/schollz/progressbar/v3"
)

func Process(Host, Title string, Page, Quantity int, writer writer.IWriter) error {
	scr := scrapper.ScrapAll(Host, Title, Page, Quantity)
	Klausa := func(objPointer *interface{}) error {
		ch := (*objPointer).(*[]model.Chapter)

		for i := 0; i < len(*ch); {
			var syncG sync.WaitGroup
			chpter := (*ch)[i]
			errChannel := make(chan error)
			errorSignal := make(chan bool)
			Images := chpter.ReconstructImage()
			syncG.Add(len(Images))
			bar := progressbar.Default(int64(len(Images)-1), fmt.Sprintf(`downloading chapter-%s`, chpter.Id))
			writer.OnFinished(func() {
				bar.Add(1)
			})
			for j := 0; j < len(Images); j++ {
				go func(ci model.ComicInfo, img model.Image) {
					defer syncG.Done()
					err := writer.Store(img, ci)
					if err != nil {
						errChannel <- err
					}
				}(model.ComicInfo{
					Title:   scr.ComicFlat,
					Chapter: chpter.ChapterFlat,
				}, Images[j])
			}

			go func() {
				errs := make([]error, 0)
				for {
					select {
					case err, ok := <-errChannel:
						if !ok {
							if len(errs) > 0 {
								log.Println(errs)
							}
							errorSignal <- len(errs) > 0
							return
						}
						errs = append(errs, err)
					}
				}
			}()

			syncG.Wait()
			close(errChannel)
			isError := <-errorSignal
			if !isError {
				*ch = append((*ch)[:i], (*ch)[i+1:]...)
				continue
			}
			i++
		}
		return errors.New("End")
	}
	stopper := func(objPointer *interface{}) bool {
		ch := (*objPointer).(*[]model.Chapter)
		return len(*ch) < 1
	}
	wrapper := retry.RecurseTry(Klausa, stopper, 3, time.Duration(2*time.Second))
	intfPointer := new(interface{})
	*intfPointer = &scr.Chapters
	err := wrapper(intfPointer)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
