package main

import (
	"log"
	"net"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("static", r.URL.Path))
	})

	hub := NewHub()
	go hub.run()
	http.HandleFunc("/sync", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	lis, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("server listening at %v", lis.Addr())
	log.Panic(http.Serve(lis, nil))
}
