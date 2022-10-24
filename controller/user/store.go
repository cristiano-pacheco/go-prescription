package controller

import (
	"log"
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/model"
	"github.com/cristiano-pacheco/go-prescription/validator"
	"github.com/cristiano-pacheco/go-prescription/view"
)

type userStoreAction struct {
	userModel *model.UserModel
}

type userCreateForm struct {
	Name   string
	Errors map[string]string
}

func NewUserStoreAction(userModel *model.UserModel) *userStoreAction {
	return &userStoreAction{userModel: userModel}
}

func (action *userStoreAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	uv := validator.UserValidator{}
	validationErrors := uv.Validate(r.PostForm.Get("name"))
	form := userCreateForm{
		Name:   r.PostForm.Get("name"),
		Errors: validationErrors,
	}

	templateData := view.TemplateData{}
	templateData.Form = form

	if len(validationErrors) > 0 {
		view.NewTemplate().RenderFormWithValidationErrors(w, &templateData)
		return
	}

	_, err = action.userModel.Insert(r.PostForm.Get("name"))
	if err != nil {
		log.Printf("An error ocurriend during the user creation insert: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
