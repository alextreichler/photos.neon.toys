package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/straightbuggin/photos.neon.toys/controllers"
	"github.com/straightbuggin/photos.neon.toys/models"
	"github.com/straightbuggin/photos.neon.toys/templates"
	"github.com/straightbuggin/photos.neon.toys/views"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

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

	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"tailwind.gohtml", "navigation.gohtml", "footer.gohtml", "signin.gohtml",
	))

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")

	csrfKey := "gFvi45R4f15xNblnEeZOQbfAVCYEIAU7"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: fix
		csrf.Secure(false),
	)
	http.ListenAndServe(":3000", csrfMw(r))
}
