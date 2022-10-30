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

	// Validates the form data
	uv := validator.UserValidator{}
	validationErrors := uv.Validate(r.PostForm.Get("name"))
	form := userUpdateForm{
		Name:   r.PostForm.Get("name"),
		Errors: validationErrors,
	}
	templateData := view.TemplateData{}
	templateData.Form = form

	// If the form data is invalid, render the page with the validation errors
	if len(validationErrors) > 0 {
		view.NewTemplate().RenderFormWithValidationErrors(w, &templateData)
		return
	}

	// Updates the user in the database
	err = action.userModel.Update(uint(id), r.PostForm.Get("name"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
