package main

import (
	"log"
	"net/http"
)

func main() {
	server := &BookServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
