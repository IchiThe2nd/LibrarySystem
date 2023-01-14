package main

import (
	"log"
	"net/http"
)

type InMemoryBookStore struct{}

func (i *InMemoryBookStore) GetISBNTitle(isbn string) string {
	return "Title"
}

func main() {
	server := &BookServer{&InMemoryBookStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
