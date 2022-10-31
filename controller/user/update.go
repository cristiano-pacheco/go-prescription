package controller

import (
	"net/http"
	"strconv"

	"github.com/alexedwards/flow"
	"github.com/cristiano-pacheco/go-prescription/model"
	"github.com/cristiano-pacheco/go-prescription/validator"
	"github.com/cristiano-pacheco/go-prescription/view"
)

type userUpdateAction struct {
	userModel *model.UserModel
}

func NewUserUpdateAction(userModel *model.UserModel) *userUpdateAction {
	return &userUpdateAction{userModel: userModel}
}

func (action *userUpdateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type userUpdateForm struct {
		ID   int
		Name string
		validator.Validator
	}

	// Parse the user ID from URL
	id, err := strconv.Atoi(flow.Param(r.Context(), "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	// Parse the form
	err = r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// If the form data is invalid, render the page with the validation errors
	form := userUpdateForm{
		ID:   id,
		Name: r.PostForm.Get("name"),
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "The name field cannot be blank")

	if !form.IsValid() {
		data := view.NewTemplateData()
		data.Form = form
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusUnprocessableEntity)
		view.NewTemplate().Render(w, data, "./view/layout/bootstrap.gohtml", "./view/user/edit.gohtml")
		return
	}

	// Updates the user in the database
	err = action.userModel.Update(uint(id), r.PostForm.Get("name"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
