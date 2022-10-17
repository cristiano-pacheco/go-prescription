package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/flow"
	controller "github.com/cristiano-pacheco/go-prescription/Controller"
)

func main() {
	router := flow.New()

	router.HandleFunc("/", controller.HomeController, "GET")

	err := http.ListenAndServe(":8000", router)
	log.Fatal(err)
}
