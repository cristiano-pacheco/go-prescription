package controller

import (
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/controller"
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
	users, err := action.userModel.GetAll()
	if err != nil {
		controller.RenderServerError(w, err)
		return
	}
	view.NewTemplate().Render(w, users, "./view/layout/bootstrap.gohtml", "./view/user/index.gohtml")
}
