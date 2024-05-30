package main

import "net/http"

func main() {
	server := Server{}

	serve := server.server()

	_ = http.ListenAndServe(":8080", serve)
}
