package controller

import (
	"fmt"
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/model"
)

type userDestroyAction struct {
	userModel *model.UserModel
}

func NewUserDestroyAction(userModel *model.UserModel) *userDestroyAction {
	return &userDestroyAction{userModel: userModel}
}

func (action *userDestroyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>User Destroy Action</h1>")
}
