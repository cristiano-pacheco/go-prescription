package controller

import (
	"fmt"
	"net/http"
)

func UserDestroyAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>User Destroy Action</h1>")
}
