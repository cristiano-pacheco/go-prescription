package view

import (
	"html/template"
	"log"
	"net/http"
)

func NewTemplate() *Template {
	return &Template{}
}

type Template struct{}

type TemplateData struct {
	Form any
}

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

func (t *Template) RenderFormWithValidationErrors(w http.ResponseWriter, templateData *TemplateData) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusUnprocessableEntity)
	t.Render(w, templateData, "./view/layout/bootstrap.gohtml", "./view/user/create.gohtml")
}
