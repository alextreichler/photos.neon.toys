package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/straightbuggin/photos.neon.toys/controllers"
	"github.com/straightbuggin/photos.neon.toys/models"
	"github.com/straightbuggin/photos.neon.toys/templates"
	"github.com/straightbuggin/photos.neon.toys/views"
)

func main() {

	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "navigation.gohtml", "footer.gohtml", "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"tailwind.gohtml", "navigation.gohtml", "footer.gohtml", "signup.gohtml",
	))

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
