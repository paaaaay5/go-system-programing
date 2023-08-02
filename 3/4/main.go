package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=result.zip")
	w.Header().Set("Content-Type", "application/zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	contentWriter, err := zipWriter.Create("content.txt")
	if err != nil {
		panic(err)
	}
	contentReader := strings.NewReader("Hello")

	io.Copy(contentWriter, contentReader)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("server running")
}
