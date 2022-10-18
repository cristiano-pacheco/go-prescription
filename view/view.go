package view

import (
	"log"
	"net/http"
	"text/template"
)

func NewTemplate() *Template {
	return &Template{}
}

type Template struct{}

func (t *Template) Render(w http.ResponseWriter, data interface{}, templates ...string) {
	templateSet, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}
	err = templateSet.ExecuteTemplate(w, "bootstrap", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}
