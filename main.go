package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/alexedwards/flow"
	homeController "github.com/cristiano-pacheco/go-prescription/controller/home"
	userController "github.com/cristiano-pacheco/go-prescription/controller/user"
)

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	router := flow.New()

	router.HandleFunc("/", homeController.HomeIndexAction, "GET")

	router.HandleFunc("/users", userController.UserIndexAction, "GET")
	router.HandleFunc("/users/create", userController.UserCreateAction, "GET")
	router.HandleFunc("/users", userController.UserStoreAction, "POST")
	router.HandleFunc("/users/:id/edit", userController.UserEditAction, "GET")
	router.HandleFunc("/users/:id/update", userController.UserUpdateAction, "POST")
	router.HandleFunc("/users/:id/destroy", userController.UserDestroyAction, "POST")

	err := http.ListenAndServe(*addr, router)

	log.Fatal(err)
}
