package main

import "net/http"

func HandleRoutes() {
	http.HandleFunc("/", home)
}
