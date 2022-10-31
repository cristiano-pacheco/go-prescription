package controller

import (
	"net/http"
	"strconv"

	"github.com/alexedwards/flow"
	"github.com/cristiano-pacheco/go-prescription/model"
	"github.com/cristiano-pacheco/go-prescription/validator"
	"github.com/cristiano-pacheco/go-prescription/view"
)

type userEditAction struct {
	userModel *model.UserModel
}

func NewUserEditAction(userModel *model.UserModel) *userEditAction {
	return &userEditAction{userModel: userModel}
}

func (action *userEditAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type userUpdateForm struct {
		ID   int
		Name string
		validator.Validator
	}
	w.Header().Set("Content-Type", "text/html")
	id, err := strconv.Atoi(flow.Param(r.Context(), "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	user, err := action.userModel.Get(uint(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	form := userUpdateForm{
		ID:   user.ID,
		Name: user.Name,
	}
	data := view.NewTemplateData()
	data.Form = form
	view.NewTemplate().Render(w, data, "./view/layout/bootstrap.gohtml", "./view/user/edit.gohtml")
}
