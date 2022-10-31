package controller

import (
	"net/http"
	"strconv"

	"github.com/cristiano-pacheco/go-prescription/controller"
	"github.com/cristiano-pacheco/go-prescription/model"
)

type userDestroyAction struct {
	userModel *model.UserModel
}

func NewUserDestroyAction(userModel *model.UserModel) *userDestroyAction {
	return &userDestroyAction{userModel: userModel}
}

func (action *userDestroyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(controller.ReadParameter(r, "id"))
	if err != nil {
		controller.RenderServerError(w, err)
		return
	}
	err = action.userModel.Delete(uint(id))
	if err != nil {
		controller.RenderServerError(w, err)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
