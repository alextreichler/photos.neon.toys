package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/straightbuggin/photos.neon.toys/controllers"
	"github.com/straightbuggin/photos.neon.toys/templates"
	"github.com/straightbuggin/photos.neon.toys/views"
)

func main() {

	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "navigation.gohtml", "footer.gohtml", "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "navigation.gohtml", "footer.gohtml", "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
