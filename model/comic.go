package model


type Comic struct {
	ComicFlat
	Chapters []Chapter
}
type ComicFlat struct {
	Id int
	Name string
	Host string
}

type Chapter struct {
	ChapterFlat
	Images []Image
}
type ChapterFlat struct {
	Id string
}

type ComicInfo struct {
	Title ComicFlat
	Chapter ChapterFlat
}

type Image struct {
	Link map[string]string
	Episode int
}