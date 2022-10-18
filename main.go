package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/alexedwards/flow"
	homeController "github.com/cristiano-pacheco/go-prescription/controller/home"
	userController "github.com/cristiano-pacheco/go-prescription/controller/user"
	"github.com/cristiano-pacheco/go-prescription/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	dsn := flag.String("dsn", "root:root@/prescription?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Database initialization
	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Router inicialization
	router := flow.New()

	// Models inicialization
	userModel := model.NewUserModel(db)

	// Actions inicialization
	userActions := userController.CreateUserActions(userModel)

	router.HandleFunc("/", homeController.HomeIndexAction, "GET")

	router.Handle("/users", userActions.IndexAction, "GET")
	router.Handle("/users/create", userActions.CreateAction, "GET")
	router.Handle("/users", userActions.StoreAction, "POST")
	router.Handle("/users/:id/edit", userActions.EditAction, "GET")
	router.Handle("/users/:id/update", userActions.UpdateAction, "POST")
	router.Handle("/users/:id/destroy", userActions.DestroyAction, "POST")

	log.Printf("Starting server on %s", *addr)
	err = http.ListenAndServe(*addr, router)
	log.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
