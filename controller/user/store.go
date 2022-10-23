package controller

import (
	"log"
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/model"
)

type userStoreAction struct {
	userModel *model.UserModel
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
	_, err = action.userModel.Insert(r.PostForm.Get("name"))
	if err != nil {
		log.Printf("An error ocurriend during the user creation insert: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
