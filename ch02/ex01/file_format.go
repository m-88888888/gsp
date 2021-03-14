package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	err := WriteFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func WriteFile(filename string) error {
	pos := strings.LastIndex(filename, ".")
	if pos == -1 {
		return errors.New("filename needs .txt")
	}

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "1+1=%d\n%s\n%f\n", 1+1, "test", 3.14)
	return nil
}
