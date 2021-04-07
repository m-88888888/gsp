package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("test.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("hoge.txt")
	if err != nil {
		panic(err)
	}

	reader := strings.NewReader("content of file in zip")
	io.Copy(writer, reader)
}
