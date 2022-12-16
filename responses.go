package main

type genericResponse struct {
	Data interface{} `json:"data"`
}

type notFoundResponse struct {
	Message string `json:"message"`
}
