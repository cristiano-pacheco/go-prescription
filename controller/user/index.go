package controller

import (
	"log"
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/model"
	"github.com/cristiano-pacheco/go-prescription/view"
)

type userIndexAction struct {
	userModel *model.UserModel
}

func NewUserIndexAction(userModel *model.UserModel) *userIndexAction {
	return &userIndexAction{userModel: userModel}
}

func (action *userIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	users, err := action.userModel.GetAll()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	view.NewTemplate().Render(w, users, "./view/layout/bootstrap.gohtml", "./view/user/index.gohtml")
}
