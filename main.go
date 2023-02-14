package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	hub := NewHub()
	go hub.run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqFilePath := filepath.Join("static", r.URL.Path)
		if reqFilePath == "static" {
			reqFilePath = "static/index.html"
		}
		log.Println(reqFilePath)
		reqFileBytes, err := os.ReadFile(reqFilePath)
		if err != nil {
			log.Printf("error: %v", err)
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}
		if reqFilePath == "static/index.html" {
			reqFileText := string(reqFileBytes)
			reqFileText = strings.ReplaceAll(reqFileText, "{{}}", hub.getCurrText())
			reqFileBytes = []byte(reqFileText)
			w.Write(reqFileBytes)
		} else {
			http.ServeFile(w, r, filepath.Join("static", r.URL.Path))
		}
	})

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
