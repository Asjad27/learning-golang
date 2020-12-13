package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("File name must be provided as an argument")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Unable to open file for reading. Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
