package main

import (
	"io"
	"os"
	"strings"
)

func myCopyN(dest io.Writer, src io.Reader, length int) {
	newReader := io.LimitReader(src, int64(length))
	io.Copy(dest, newReader)
}

func main() {
	writer := os.Stdout
	reader := strings.NewReader("123456789")
	myCopyN(writer, reader, 4)
}
