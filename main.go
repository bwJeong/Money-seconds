package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed index.html
var webFS embed.FS

func main() {
	http.Handle("/", http.FileServerFS(webFS))
	addr := ":8080"
	log.Printf("listening on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
