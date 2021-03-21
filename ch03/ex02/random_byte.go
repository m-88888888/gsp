package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "tmpfile.txt"
	err := RandomByte(fileName)
	if err != nil {
		panic(err)
	}
	got, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(got))

}

func RandomByte(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	io.CopyN(file, rand.Reader, 1024)
	return nil
}
