package transport

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/post", post).Methods("POST")
	m.HandleFunc("/get", get).Methods("GET")
	return m
}
