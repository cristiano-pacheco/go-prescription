package controller

import (
	"fmt"
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
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>User Store Action</h1>")
}
