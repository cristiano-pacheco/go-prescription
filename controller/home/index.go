package controller

import (
	"fmt"
	"net/http"
)

func HomeIndexAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Oh Yeah!</h1>")
}
