package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader
	charA := io.NewSectionReader(programming, 5, 1)
	charS := io.NewSectionReader(system, 0, 1)
	charC := io.NewSectionReader(computer, 0, 1)
	charI1 := io.NewSectionReader(programming, 8, 1)
	charI2 := io.NewSectionReader(programming, 8, 1)
	stream = io.MultiReader(charA, charS, charC, charI1, charI2)

	io.Copy(os.Stdout, stream)
}
