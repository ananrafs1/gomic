package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/ananrafs1/gomic/lib/monitoring"
	"github.com/ananrafs1/gomic/model"
	handlerdownload "github.com/ananrafs1/gomic/pkg/server/handlerDownload"
	handlerstatic "github.com/ananrafs1/gomic/pkg/server/handlerStatic"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	monitoring.Initialize()

	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(handlerstatic.IndexHandler()), "index"))

	for title, files := range handlerstatic.GenerateStaticContentURLMapping() {
		fileServerPattern := fmt.Sprintf(`%s/img/%s/`, model.Config.Server["readRoot"], title)
		http.Handle(fileServerPattern,
			http.StripPrefix(fileServerPattern,
				http.FileServer(http.Dir(filepath.Join(model.Config.OutputDir, title)))))
		http.Handle(fmt.Sprintf(`%s/%s`, model.Config.Server["readRoot"], title), otelhttp.NewHandler(http.HandlerFunc(handlerstatic.StaticFileHandler(files, title)), title))
	}

	http.Handle(model.Config.Server["downloadRoot"], otelhttp.NewHandler(http.HandlerFunc(handlerdownload.DownloaderHandler()), "downloader"))

	fmt.Println(http.ListenAndServe(fmt.Sprintf(`:%s`, model.Config.Server["port"]), nil))

}
func init() {
	model.InitAndListenConfig()
}
