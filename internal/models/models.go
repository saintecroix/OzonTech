package models

var Inmemory = make(map[string]string)

type Urls struct {
	Orig  string `json:"orig"`
	Short string `json:"shortname"`
}
