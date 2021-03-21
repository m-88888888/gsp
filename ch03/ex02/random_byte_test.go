package main

import (
	"io/ioutil"
	"testing"
)

func TestRandomByte(t *testing.T) {
	fileName := "tmpfile_test.txt"
	err := RandomByte(fileName)
	if err != nil {
		t.Fatal(err)
	}
	got, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1024 {
		t.Errorf("result = %v, expected = %v", len(got), 1024)
	}
}
