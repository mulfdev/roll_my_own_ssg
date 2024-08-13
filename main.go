package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")

	fileInfo, err := os.Stat("./markdown")
	if err != nil {
		panic("markdown directory not found")
	}

	isDir := fileInfo.IsDir()

	if isDir != true {
		panic("Must be a directory")
	}

}
