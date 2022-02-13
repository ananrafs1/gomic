package core

import (
	"github.com/ananrafs1/gomic/scrapper"
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/writer"
	"github.com/ananrafs1/go-palugada/retry"
	"time"
	"sync"
	"log"
	"fmt"
	"github.com/schollz/progressbar/v3"
)

func Proces(Host, Title string, writer writer.IWriter ) error {
	scr, err := scrapper.ScrapAll(Host, Title)
	log.Println("1", len(scr.Chapters), err)
	Klausa := func(objPointer *interface{}) error {
		ch := (*objPointer).(*[]model.Chapter)
		log.Println(ch)


		for i:= 1; i < len(*ch);  {
			var syncG sync.WaitGroup
			chpter := (*ch)[i]
			errChannel := make(chan error)
			errorSignal := make(chan bool)
			syncG.Add(len(chpter.Images))
			bar := progressbar.Default(int64(len(chpter.Images)-1), fmt.Sprintf(`downloading chapter-%s`, chpter.Id))
			writer.OnFinished(func(){
				bar.Add(1)
			})
			for j:= 0; j < len(chpter.Images); j++ {
				go func(ci model.ComicInfo, img model.Image) {
					defer syncG.Done()
					err := writer.Store(img, ci)
					if err != nil {
						errChannel <- err
					}
				}(model.ComicInfo{
					scr.ComicFlat,
					chpter.ChapterFlat,
				}, chpter.Images[j])
			}

			go func(){
				errs := make([]error,0)
				for {
					select {
					case err, ok := <- errChannel:
						if !ok {
							log.Println("error channel closed")
							errorSignal <-  len(errs) > 0
						}
						errs = append(errs, err)
					}
				}
			}()

			syncG.Wait()
			close(errChannel)
			isError := <- errorSignal
			if !isError {
				*ch = append((*ch)[:i], (*ch)[i+1:]...)
				continue
			}
			i++
		}
		return nil
	}
	stopper := func (objPointer *interface{}) bool {
		ch := (*objPointer).(*[]model.Chapter)
		return len(*ch) < 1
	}
	wrapper := retry.RecurseTry(Klausa, stopper, 3, time.Duration(2*time.Second))
	intfPointer := new(interface{})
	*intfPointer = &scr.Chapters
	err = wrapper(intfPointer)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}