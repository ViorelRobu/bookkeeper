package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func notFound(w http.ResponseWriter, r *http.Request) {
	message := notFoundResponse{
		Message: "404 page not found",
	}

	res, err := json.Marshal(message)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, string(res))
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	message := genericResponse{
		Data: getAllBooks().Result,
	}

	res, err := json.Marshal(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, string(res))
}
