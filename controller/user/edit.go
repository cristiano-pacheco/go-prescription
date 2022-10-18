package controller

import (
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/model"
	"github.com/cristiano-pacheco/go-prescription/view"
)

type userEditAction struct {
	userModel *model.UserModel
}

func NewUserEditAction(userModel *model.UserModel) *userEditAction {
	return &userEditAction{userModel: userModel}
}

func (action *userEditAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	view.NewTemplate().Render(w, nil, "./view/layout/bootstrap.gohtml", "./view/user/edit.gohtml")
}
