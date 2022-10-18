package controller

import (
	"fmt"
	"net/http"
)

func UserStoreAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>User Store Action</h1>")
}
