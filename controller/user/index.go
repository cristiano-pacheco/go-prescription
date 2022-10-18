package controller

import (
	"log"
	"net/http"
	"text/template"
)

func UserIndexAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	files := []string{"./view/layout/bootstrap.gohtml", "./view/user/index.gohtml"}
	templateSet, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}
	err = templateSet.ExecuteTemplate(w, "bootstrap", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}
