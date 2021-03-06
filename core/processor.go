package core

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ananrafs1/go-palugada/retry"
	"github.com/ananrafs1/gomic/lib/monitoring"
	"github.com/ananrafs1/gomic/lib/scrapper"
	"github.com/ananrafs1/gomic/lib/writer"
	"github.com/ananrafs1/gomic/lib/writer/interceptor"
	"github.com/ananrafs1/gomic/model"
	"github.com/schollz/progressbar/v3"
)

func Process(ctx context.Context, Host, Title string, Page, Quantity int, writer writer.IWriter) error {
	_, closer, ctx := monitoring.CreateSpan(ctx, "Process")
	defer closer()
	scr := scrapper.ScrapAll(ctx, Host, Title, Page, Quantity)
	return Store(ctx, scr, writer)
}

func Store(ctx context.Context, scr *model.Comic, writer writer.IWriter) error {
	_, closer, ctx := monitoring.CreateSpan(ctx, "ProcessStore")
	defer closer()
	wrapper := retry.RecurseTry(coreProcess(ctx, writer, scr),
		stopperProcess,
		3,
		time.Duration(2*time.Second))
	intfPointer := new(interface{})
	*intfPointer = &scr.Chapters
	err := wrapper(intfPointer)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func coreProcess(ctx context.Context, writer writer.IWriter, scr *model.Comic) func(objPointer *interface{}) error {
	return func(objPointer *interface{}) error {
		ch := (*objPointer).(*[]model.Chapter)

		for i := 0; i < len(*ch); {
			var syncG sync.WaitGroup
			errChannel, errorSignal := make(chan error), make(chan bool)

			chpter := (*ch)[i]
			Images := chpter.ReconstructImage()

			bar := progressbar.Default(int64(len(Images)-1), fmt.Sprintf(`downloading chapter-%s`, chpter.Id))
			writerInterceptor := interceptor.WriterProcess{WR: writer}
			writerInterceptor.OnFinished(func() { bar.Add(1) })
			writer := interceptor.WriterMonitor{WR: &writerInterceptor}

			syncG.Add(len(Images))
			for j := 0; j < len(Images); j++ {
				go func(ci model.ComicInfo, img model.Image) {
					defer syncG.Done()
					err := writer.Store(ctx, img, ci)
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
}

func stopperProcess(objPointer *interface{}) bool {
	ch := (*objPointer).(*[]model.Chapter)
	return len(*ch) < 1
}
