package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}

	writer := gzip.NewWriter(w)
	writers := io.MultiWriter(os.Stdout, writer)

	encoder := json.NewEncoder(writers)
	encoder.SetIndent("", "　")
	encoder.Encode(source)
}
