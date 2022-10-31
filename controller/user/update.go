package controller

import (
	"net/http"
	"strconv"

	"github.com/cristiano-pacheco/go-prescription/controller"
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
	id, err := strconv.Atoi(controller.ReadParameter(r, "id"))
	if err != nil {
		controller.RenderServerError(w, err)
		return
	}

	// Parse the form
	err = r.ParseForm()
	if err != nil {
		controller.RenderServerError(w, err)
		return
	}

	// If the form data is invalid, render the page with the validation errors
	form := userUpdateForm{
		ID:   id,
		Name: r.PostForm.Get("name"),
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "The name field cannot be blank")
	form.CheckField(validator.MaxChars(form.Name, 255), "name", "The name field cannot be more than 255 characters long")
	form.CheckField(validator.GreaterThanZero(form.ID), "id", "The given user has to be greater than zero")

	if !form.IsValid() {
		data := view.NewTemplateData()
		data.Form = form
		w.WriteHeader(http.StatusUnprocessableEntity)
		view.NewTemplate().Render(w, data, "./view/layout/bootstrap.gohtml", "./view/user/edit.gohtml")
		return
	}

	// Updates the user in the database
	err = action.userModel.Update(uint(id), r.PostForm.Get("name"))
	if err != nil {
		controller.RenderServerError(w, err)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
