package router

import (
	"github.com/alexedwards/flow"
	homeController "github.com/cristiano-pacheco/go-prescription/controller/home"
	userActionsFactory "github.com/cristiano-pacheco/go-prescription/controller/user"
	"github.com/cristiano-pacheco/go-prescription/model"
)

type Router struct {
	Mux *flow.Mux
}

func CreateRouter(models *model.Models) *Router {
	mux := flow.New()

	userActions := userActionsFactory.CreateUserActions(models.UserModel)

	mux.HandleFunc("/", homeController.HomeIndexAction, "GET")

	mux.Handle("/users", userActions.IndexAction, "GET")
	mux.Handle("/users/create", userActions.CreateAction, "GET")
	mux.Handle("/users/store", userActions.StoreAction, "POST")
	mux.Handle("/users/:id/edit", userActions.EditAction, "GET")
	mux.Handle("/users/:id/update", userActions.UpdateAction, "POST")
	mux.Handle("/users/:id/destroy", userActions.DestroyAction, "POST")

	return &Router{Mux: mux}
}
