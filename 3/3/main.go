package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("./3.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("./newFile.txt")
	if err != nil {
		panic(err)
	}

	s := "Hello"
	stringReader := strings.NewReader(s)
	io.Copy(writer, stringReader)

}
