package controller

import (
	"net/http"
	"strconv"

	"github.com/alexedwards/flow"
	"github.com/cristiano-pacheco/go-prescription/model"
)

type userDestroyAction struct {
	userModel *model.UserModel
}

func NewUserDestroyAction(userModel *model.UserModel) *userDestroyAction {
	return &userDestroyAction{userModel: userModel}
}

func (action *userDestroyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(flow.Param(r.Context(), "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	err = action.userModel.Delete(uint(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
