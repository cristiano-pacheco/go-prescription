package controller

import (
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/view"
)

func UserIndexAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	view.NewTemplate().Render(w, nil, "./view/layout/bootstrap.gohtml", "./view/user/index.gohtml")
}
