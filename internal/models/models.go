package models

var Inmemory = make(map[string]string)

type Urls struct {
	Orig  string `json:"url"`
	Short string `json:"short-url"`
}
