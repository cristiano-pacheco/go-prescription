package controller

import (
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/view"
)

func HomeIndexAction(w http.ResponseWriter, r *http.Request) {
	view.NewTemplate().Render(w, nil, "./view/layout/bootstrap.gohtml", "./view/home/index.gohtml")
}
