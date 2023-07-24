package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/straightbuggin/photos.neon.toys/controllers"
	"github.com/straightbuggin/photos.neon.toys/migrations"
	"github.com/straightbuggin/photos.neon.toys/models"
	"github.com/straightbuggin/photos.neon.toys/templates"
	"github.com/straightbuggin/photos.neon.toys/views"
)

func main() {
	// Setup Database
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// Setup Services
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	// Setup Middleware
	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}
	csrfKey := "gFvi45R4f15xNblnEeZOQbfAVCYEIAU7"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: fix
		csrf.Secure(false),
	)

	// Setup controllers
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"tailwind.gohtml", "navigation.gohtml", "footer.gohtml", "signup.gohtml",
	))

	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"tailwind.gohtml", "navigation.gohtml", "footer.gohtml", "signin.gohtml",
	))

	// Setup Routes
	r := chi.NewRouter()
	r.Use(csrfMw)
	r.Use(umw.SetUser)
	r.Use(middleware.Logger)

	tpl := views.Must(views.ParseFS(templates.FS,
		"tailwind.gohtml",
		"navigation.gohtml",
		"footer.gohtml",
		"home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})
	//	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
