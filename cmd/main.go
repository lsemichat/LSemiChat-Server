package main

import "net/http"

func main() {
	http.HandleFunc("/ping", pingHandler)

	http.ListenAndServe(":8080", nil)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
