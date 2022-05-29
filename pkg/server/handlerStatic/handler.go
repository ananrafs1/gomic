package handlerstatic

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/ananrafs1/gomic/lib/monitoring"
	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/utils"
	"go.opentelemetry.io/otel/label"
)

func IndexHandler() func(http.ResponseWriter, *http.Request) {
	mapURL := GenerateStaticContentURLMapping()
	urls := make([]string, 0)
	for title, _ := range mapURL {
		urls = append(urls, fmt.Sprintf("%s/%s", model.Config.Server["readRoot"], title))
	}
	return baseIndexHandler(urls)
}

func baseIndexHandler(urls []string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		span, closer, ctx := monitoring.CreateSpan(ctx, "baseIndexHandler")
		span.SetAttributes(label.String("urls", strings.Join(urls, ",")))
		defer closer()
		var err error
		defer func() {
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("403 - Forbidden"))
			}
			// w.WriteHeader(http.StatusOK)
			// w.Write([]byte(strings.Join(urls, ", ")))
		}()

		templateDir := filepath.Join(utils.GetRootDir(), model.Config.Server["staticDir"], "index.html")
		templates, err := template.ParseFiles(templateDir)
		if err != nil {
			panic(err)
		}
		templates.Execute(w, urls)
	}
}

func StaticFileHandler(chapters map[string][]string, title string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := r.Context()
		span, closer, ctx := monitoring.CreateSpan(ctx, "StaticFileHandler")

		defer closer()
		templateDir := filepath.Join(utils.GetRootDir(),
			model.Config.Server["staticDir"],
			"template.html")
		templates, err := template.ParseFiles(templateDir)
		if err != nil {
			panic(err)
		}
		chapterQS := "1"
		if len(r.URL.Query()["ch"]) > 0 {
			chapterQS = r.URL.Query()["ch"][0]
		}
		nt := struct {
			Img  []string
			Next string
			Prev string
		}{}

		for _, v := range chapters[chapterQS] {
			nt.Img = append(nt.Img, fmt.Sprintf("img/%s/%s/%s", title, chapterQS, v))
		}

		intChapter, err := strconv.Atoi(chapterQS)
		if err != nil {
			panic(err)
		}
		nt.Next = fmt.Sprintf(`%s?ch=%d`, title, intChapter+1)
		nt.Prev = fmt.Sprintf(`%s?ch=%d`, title, intChapter-1)
		templates.Execute(w, nt)
		span.SetAttributes(label.String("chapter", chapterQS), label.String("title", title))
	}
}

func GenerateStaticContentURLMapping() map[string]map[string][]string {
	targetpath := filepath.Join(model.Config.OutputDir)
	items, _ := ioutil.ReadDir(targetpath)
	MapChapterFiles := map[string]map[string][]string{}
	for _, item := range items {
		title, chapters := getTitle(item)
		for k := range chapters {
			sort.SliceStable(chapters[k], func(i, j int) bool {
				convertToDigit := func(val string) (ret int, canConvert bool) {
					noExt := strings.Split(val, ".")[0]
					ret, err := strconv.Atoi(noExt)
					if err != nil {
						return ret, false
					}
					canConvert = true
					return
				}
				leftSide, canConvert := convertToDigit(chapters[k][i])
				if !canConvert {
					return false
				}
				rightSide, canConvert := convertToDigit(chapters[k][j])
				if !canConvert {
					return true
				}
				return leftSide < rightSide
			})
		}
		MapChapterFiles[title] = chapters
	}
	return MapChapterFiles
}

func getTitle(item fs.FileInfo) (title string, chapters map[string][]string) {
	if item.IsDir() {
		title = item.Name()
		chapters = make(map[string][]string)
		getFiles(filepath.Join(model.Config.OutputDir, title), chapters)
	}
	return
}

func getFiles(currentPath string, chapters map[string][]string) {
	subitems, _ := os.ReadDir(currentPath)
	for _, subitem := range subitems {
		fileName := subitem.Name()
		if subitem.IsDir() {
			getFiles(filepath.Join(currentPath, fileName), chapters)
		} else {
			currentDir, err := os.Stat(currentPath)
			if err != nil {
				panic(err)
			}
			dirName := currentDir.Name()
			val, ok := chapters[dirName]
			if !ok {
				chapters[dirName] = []string{fileName}
				continue
			}
			val = append(val, fileName)
			chapters[dirName] = val
		}
	}
}
