package main

import (
	"io/ioutil"
	"testing"
)

func TestWriteFile(t *testing.T) {
	var tests = []struct {
		name      string
		filename  string
		expected  string
		isSucceed bool
	}{
		{
			name:      "正常系",
			filename:  "testfile.txt",
			expected:  "1+1=2\ntest\n3.140000\n",
			isSucceed: true,
		},
	}
	for _, test := range tests {
		t.Run("case Normal", func(t *testing.T) {
			err := WriteFile(test.filename)
			if err != nil {
				t.Fatalf("%v\n", err)
			}
			got, err := ioutil.ReadFile(test.filename)
			if err != nil {
				t.Fatalf("%v\n", err)
			}
			if string(got) != test.expected {
				t.Errorf("%q, result = %q\n, expected = %q", test.name, got, test.expected)
			}
		})
	}

}
