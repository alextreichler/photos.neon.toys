package controllers

import (
	"net/http"

	"github.com/straightbuggin/photos.neon.toys/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
