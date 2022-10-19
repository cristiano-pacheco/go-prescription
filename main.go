package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/cristiano-pacheco/go-prescription/model"
	"github.com/cristiano-pacheco/go-prescription/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// APP flags declaration
	addr := flag.String("addr", ":8000", "HTTP network address")
	dsn := flag.String("dsn", "root:root@/prescription?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Database initialization
	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Models inicialization
	models := model.CreateModels(db)

	// Routes inicialization
	router := router.CreateRouter(models)

	log.Printf("Starting server on %s", *addr)
	err = http.ListenAndServe(*addr, router.Mux)
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
