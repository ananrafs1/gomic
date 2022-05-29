package handlerdownload

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ananrafs1/gomic/core"
	"github.com/ananrafs1/gomic/lib/monitoring"
	"github.com/ananrafs1/gomic/lib/writer/downloader"
	"go.opentelemetry.io/otel/label"
)

func DownloaderHandler() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			}
			// w.WriteHeader(http.StatusOK)
			// w.Write([]byte(strings.Join(urls, ", ")))
		}()
		ctx := r.Context()
		span, closer, ctx := monitoring.CreateSpan(ctx, "DownloaderHandler")
		defer closer()
		mapQS := map[string]string{
			"p": "",
			"q": "",
			"t": "",
			"h": "",
		}
		for k := range mapQS {
			temp := r.URL.Query()[k]
			if len(temp) < 1 {
				err = errors.New(fmt.Sprintf("Error, Not found : %s", k))
				break
			}
			mapQS[k] = temp[0]
		}
		if err != nil {
			return
		}
		p, err := strconv.Atoi(mapQS["p"])
		q, err := strconv.Atoi(mapQS["q"])

		err = core.Process(ctx, mapQS["h"], mapQS["t"], p, q, &downloader.Downloader{})
		span.SetAttributes(label.String("Title", mapQS["t"]), label.Int("Page", p), label.Int("Quantity", q))
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("DONE"))
	}
}
