package model


type Comic struct {
	ComicFlat 				`json:"ComicFlat"`
	Chapters []Chapter 		`json:"Chapters"`
}
type ComicFlat struct {
	Id int 					`json:"Id"`
	Name string 			`json:"Name"`
	Host string 			`json:"Host"`
}

type Chapter struct {
	ChapterFlat 			`json:"ChapterFlat"`
	Images []ImageProvider 	`json:"Images"`
}
type ChapterFlat struct {
	Id string 				`json:"Id"`
}

type ComicInfo struct {
	Title ComicFlat   		`json:"Title"`
	Chapter ChapterFlat  	`json:"Chapter"`
}

type Image struct {
	Link map[string]string  	
	Episode int
}

type ImageProvider struct {
	Episode int 			`json:"Episode"`
	Provider string 		`json:"Provider"`
	Link string 			`json:"Link"`
}

func (ch *Chapter) FlattenImage(episode int) Image {
	ret := Image{ Episode : episode}
	for i :=0; i < len(ch.Images); i++ {
		if ch.Images[i].Episode == episode {
			provider := ch.Images[i].Provider
			_, ok := ret.Link[provider]
			if !ok {
				ret.Link[provider] = ch.Images[i].Link
			}
		}
	}
	return ret
}

func (ch *Chapter) ReconstructImage() []Image {
	ret := make([]Image, 0)
	mapEpisodeIndex := map[int]int{}

	for i :=0; i < len(ch.Images); i++ {
		currentEpisode := ch.Images[i].Episode
		currentImage := Image{Episode : currentEpisode, Link: map[string]string{}}
		index, ok := mapEpisodeIndex[currentEpisode]
		if !ok {
			ret = append(ret, currentImage)
			mapEpisodeIndex[currentEpisode] = len(ret) -1
			continue
		}
		currentImage = ret[index]
		provider := ch.Images[i].Provider
		_, ok = currentImage.Link[provider]
		if !ok {
			currentImage.Link[provider] = ch.Images[i].Link
		}
	}
	return ret
}