package library

import (
	"fmt"
	"net/http"
)

func BookServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "abook")
}
