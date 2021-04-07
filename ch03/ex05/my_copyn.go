package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := "test hoge"
	buf := strings.NewReader(s)
	fmt.Printf("origin string = %v\n", s)
	n, err := myCopyN(os.Stdout, buf, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\ncopyn size = %d", n)
}

func myCopyN(dest io.Writer, src io.Reader, length int) (int, error) {
	length64 := int64(length)
	lr := io.LimitReader(src, length64)
	written, err := io.Copy(dest, lr)
	if written == length64 {
		return int(written), nil
	}
	if written < length64 && err == nil {
		return int(written), io.EOF
	}
	return int(written), err

}
