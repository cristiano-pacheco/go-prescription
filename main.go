package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/flow"
	homeController "github.com/cristiano-pacheco/go-prescription/controller/home"
	userController "github.com/cristiano-pacheco/go-prescription/controller/user"
)

func main() {
	router := flow.New()

	router.HandleFunc("/", homeController.HomeIndexAction, "GET")

	router.HandleFunc("/users", userController.UserIndexAction, "GET")
	router.HandleFunc("/users/create", userController.UserCreateAction, "GET")
	router.HandleFunc("/users", userController.UserStoreAction, "POST")
	router.HandleFunc("/users/:id/edit", userController.UserEditAction, "GET")
	router.HandleFunc("/users/:id/update", userController.UserUpdateAction, "POST")
	router.HandleFunc("/users/:id/destroy", userController.UserDestroyAction, "POST")

	err := http.ListenAndServe(":8000", router)

	log.Fatal(err)
}
