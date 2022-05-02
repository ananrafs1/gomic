package main

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

	"github.com/ananrafs1/gomic/model"
	"github.com/ananrafs1/gomic/utils"
)

func main() {
	urls := generateHandler()

	fmt.Println(urls)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})
	fmt.Println(http.ListenAndServe(fmt.Sprintf(`:%s`, model.Config.Server["port"]), nil))

}

func generateHandler() []string {
	urls := make([]string, 0)
	targetpath := filepath.Join(model.Config.OutputDir)
	items, _ := ioutil.ReadDir(targetpath)
	MapChapterFiles := map[string]map[string][]string{}
	for _, item := range items {
		title, chapters := GetTitle(item)
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
	for title, files := range MapChapterFiles {
		fileServerPattern := fmt.Sprintf(`/img/%s/`, title)
		http.Handle(fileServerPattern,
			http.StripPrefix(fileServerPattern,
				http.FileServer(http.Dir(filepath.Join(model.Config.OutputDir, title)))))
		http.HandleFunc(fmt.Sprintf(`/%s`, title), serveStaticFile(files, title))
		urls = append(urls, title)
	}
	return urls
}

func serveStaticFile(chapters map[string][]string, title string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		templateDir := filepath.Join(utils.GetRootDir(),
			model.Config.Server["staticDir"],
			"template.html")
		templates, err := template.ParseFiles(templateDir)
		if err != nil {
			panic(err)
		}
		chapterQS := r.URL.Query()["ch"]
		nt := struct {
			Img  []string
			Next string
		}{}

		for _, v := range chapters[chapterQS[0]] {
			nt.Img = append(nt.Img, fmt.Sprintf("img/%s/%s/%s", title, chapterQS[0], v))
		}

		intChapter, err := strconv.Atoi(chapterQS[0])
		if err != nil {
			panic(err)
		}
		nt.Next = fmt.Sprintf(`%s?ch=%d`, title, intChapter+1)
		templates.Execute(w, nt)
	}
}

func GetTitle(item fs.FileInfo) (title string, chapters map[string][]string) {
	if item.IsDir() {
		title = item.Name()
		chapters = make(map[string][]string)
		GetFiles(filepath.Join(model.Config.OutputDir, title), chapters)
	}
	return
}

func GetFiles(currentPath string, chapters map[string][]string) {
	subitems, _ := os.ReadDir(currentPath)
	for _, subitem := range subitems {
		if subitem.IsDir() {
			GetFiles(filepath.Join(currentPath, subitem.Name()), chapters)
		} else {
			currentDir, err := os.Stat(currentPath)
			if err != nil {
				panic(err)
			}
			fileName := subitem.Name()
			val, ok := chapters[currentDir.Name()]
			if !ok {
				chapters[currentDir.Name()] = []string{fileName}
				continue
			}
			val = append(val, fileName)
			chapters[currentDir.Name()] = val
		}
	}
}

func init() {
	model.InitAndListenConfig()
}
