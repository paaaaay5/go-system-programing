package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("2.1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%s-%s-%s", "a", "b", "c")
}
