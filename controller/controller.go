package controller

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/alexedwards/flow"
)

func ReadParameter(r *http.Request, parameterName string) string {
	return flow.Param(r.Context(), parameterName)
}

func RenderServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	log.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func RenderNotFound(w http.ResponseWriter) {
	RenderClientError(w, http.StatusNotFound)
}

func RenderClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
