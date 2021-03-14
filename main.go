package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "   ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
	res, _ := json.Marshal(encoder)
	w.Write(res)
	// io.WriteString(w, encoder)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
