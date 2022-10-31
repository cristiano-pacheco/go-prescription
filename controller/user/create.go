package controller

import (
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/model"
	"github.com/cristiano-pacheco/go-prescription/view"
)

type userCreateAction struct {
}

func NewUserCreateAction(userModel *model.UserModel) *userCreateAction {
	return &userCreateAction{}
}

func (action *userCreateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	view.NewTemplate().Render(w, nil, "./view/layout/bootstrap.gohtml", "./view/user/create.gohtml")
}
