package controller

import (
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/view"
)

func UserCreateAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	view.NewTemplate().Render(w, nil, "./view/layout/bootstrap.gohtml", "./view/user/create.gohtml")
}
