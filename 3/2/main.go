package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	randReader := rand.Reader
	file, err := os.Create("randombyte")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	r := io.LimitReader(randReader, 1024)
	io.Copy(file, r)
}
