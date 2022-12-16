package main

import (
	"net/http"
)

func HandleRoutes() {
	// all not found routes will go here, together with "/"
	http.HandleFunc("/", notFound)

	// the other routes
	http.HandleFunc("/books/all", listBooks)
}
