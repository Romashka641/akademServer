package main

import (
	"io"
	"net/http"
	"os"
)

func routes() {
	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r)
		if(err!=nil) { println(err) }
		os.WriteFile("data/data.json", []byte(b), 0666)
	})
}
