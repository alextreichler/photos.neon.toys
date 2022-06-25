package controllers

import (
	"net/http"

	"github.com/straightbuggin/photos.neon.toys/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
