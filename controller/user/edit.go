package controller

import (
	"fmt"
	"net/http"
)

func UserEditAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>User Edit Action</h1>")
}
