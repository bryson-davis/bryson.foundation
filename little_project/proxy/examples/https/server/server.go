package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	cfg := &tls.Config{}
	srv := &http.Server{
		Addr:      ":8443",
		Handler:   &handler{},
		TLSConfig: cfg,
	}
	log.Fatal(srv.ListenAndServeTLS("./examples/https/server/server.crt", "./examples/https/server/server.key"))
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("PONG\n"))
}