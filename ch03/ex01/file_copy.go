package main

import (
	"flag"
	"io"
	"os"
)

var flagvar string

func init() {
	flag.StringVar(&flagvar, "s", "new.txt", "enter the copy file name")
}

func main() {
	flag.Parse()

	file, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create(flagvar)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(newFile, file)
}
