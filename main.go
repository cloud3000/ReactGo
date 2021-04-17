package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed build

var f embed.FS

func main() {
	fsys := fs.FS(f)
	contentStatic, err := fs.Sub(fsys, "build")
	fs := http.FileServer(http.FS(contentStatic))
	http.Handle("/", fs)
	log.Println("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
