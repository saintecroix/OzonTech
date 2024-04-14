package main

import (
	"OzonTech/internal/transport"
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	flag.Parse()

	srv := &http.Server{
		Addr:    *addr,
		Handler: transport.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
