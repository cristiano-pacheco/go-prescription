package controller

import (
	"fmt"
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/model"
)

type userUpdateAction struct {
	userModel *model.UserModel
}

func NewUserUpdateAction(userModel *model.UserModel) *userUpdateAction {
	return &userUpdateAction{userModel: userModel}
}

func (action *userUpdateAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>User Update Action</h1>")
}
