package main

import (
	"fmt"
	"os"
)

func main() {
	WriteFile("test.txt")
}

func WriteFile(filename string) {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "1+1=%d\n%s\n%f\n", 1+1, "test", 3.14)
}
