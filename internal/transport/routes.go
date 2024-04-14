package transport

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/post/{url}", post).Methods("POST")
	m.HandleFunc("/get/{url}", get).Methods("GET")
	return m
}
