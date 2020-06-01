package main

import "net/http"

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Golang mycellar API"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	httpServer := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	httpServer.ListenAndServe()
}
