package main

import (
	"fmt"
	"net/http"
	"strings"
)

type BookStore interface {
	GetISBNTitle(isbn string) string
}

type BookServer struct {
	store BookStore
}

func (b *BookServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isbn := strings.TrimPrefix(r.URL.Path, "/books/")
	fmt.Fprint(w, b.store.GetISBNTitle(isbn))
}

func GetISBNTitle(isbn string) string {
	if isbn == "1234" {
		return "Title1"
	}

	return "0"
}
