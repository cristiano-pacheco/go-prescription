package controller

import (
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/model"
	"github.com/cristiano-pacheco/go-prescription/view"
)

type userIndexAction struct {
	userModel *model.UserModel
}

func NewUserIndexAction(userModel *model.UserModel) *userIndexAction {
	return &userIndexAction{userModel: userModel}
}

func (action *userIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	view.NewTemplate().Render(w, nil, "./view/layout/bootstrap.gohtml", "./view/user/index.gohtml")
}
