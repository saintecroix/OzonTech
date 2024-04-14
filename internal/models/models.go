package models

var Inmemory map[string]string

type Urls struct {
	Orig  string `json:"orig"`
	Short string `json:"short"`
}
