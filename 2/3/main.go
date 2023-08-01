package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	//json化する前のデータ
	source := map[string]string{
		"Hello": "World",
	}

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	mw := io.MultiWriter(gzipWriter, os.Stdout)
	encoder := json.NewEncoder(mw)
	encoder.SetIndent("", "   ")
	encoder.Encode(source)
	gzipWriter.Flush()

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("server running")
}
