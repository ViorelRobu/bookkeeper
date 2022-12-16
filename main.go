package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type response struct {
	Succcess bool     `json:"success"`
	Data     []string `json:"data"`
}

func main() {
	fmt.Println("Starting the web server")

	HandleRoutes()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	message := response{
		Succcess: true,
		Data:     []string{},
	}

	res, err := json.Marshal(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, string(res))
}
