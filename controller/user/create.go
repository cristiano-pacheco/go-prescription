package controller

import (
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/model"
	"github.com/cristiano-pacheco/go-prescription/view"
)

type userCreateAction struct {
	userModel *model.UserModel
}

func NewUserCreateAction(userModel *model.UserModel) *userCreateAction {
	return &userCreateAction{userModel: userModel}
}

func (action *userCreateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	view.NewTemplate().Render(w, nil, "./view/layout/bootstrap.gohtml", "./view/user/create.gohtml")
}
