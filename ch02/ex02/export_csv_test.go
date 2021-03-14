package main

import (
	"io/ioutil"
	"testing"
)

func TestWriteCsvFile(t *testing.T) {
	var tests = []struct {
		name     string
		records  [][]string
		expected string
	}{
		{
			name: "正常系",
			records: [][]string{
				{"No", "name"},
				{"1", "hoge"},
				{"2", "fuga"},
				{"3", "hogefuga"},
			},
			expected: "No,name\n1,hoge\n2,fuga\n3,hogefuga\n",
		},
	}

	for _, test := range tests {
		err := WriteCsvFile(test.records)
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		got, err := ioutil.ReadFile("test.csv")
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		if string(got) != test.expected {
			t.Errorf("%q, result = %q\n, expected = %q", test.name, got, test.expected)
		}
	}

}
