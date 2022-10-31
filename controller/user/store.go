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

func NewUserStoreAction(userModel *model.UserModel) *userStoreAction {
	return &userStoreAction{userModel: userModel}
}

func (action *userStoreAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type userCreateForm struct {
		Name string
		validator.Validator
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	form := userCreateForm{
		Name: r.PostForm.Get("name"),
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "The name field cannot be blank")

	if !form.IsValid() {
		data := view.NewTemplateData()
		data.Form = form
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusUnprocessableEntity)
		view.NewTemplate().Render(w, data, "./view/layout/bootstrap.gohtml", "./view/user/create.gohtml")
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
