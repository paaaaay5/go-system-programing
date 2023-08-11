package main

import (
	"io"
	"os"
)

func main() {
	oldFile, err := os.Open("./old.txt")
	defer oldFile.Close()
	if err != nil {
		panic(err)
	}
	newFile, err := os.Create("./new.txt")
	defer newFile.Close()
	if err != nil {
		panic(err)
	}
	io.Copy(newFile, oldFile)
}
