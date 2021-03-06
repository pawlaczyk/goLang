package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("static/", http.StripPrefix("static", files))

	mux.HandleFunc("/")

	server := &http.Server{
		Addr:    "0.0.0.0:808",
		Handler: mux,
	}

	server.ListenAndServe()
}
