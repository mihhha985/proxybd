package main

import (
	"log"
	"net/http"

	_ "test/docs"

	"test/config"
	"test/internal/user"
	"test/pkg/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title ProxyDB User Repository
// @version 1.0
// @description Simple CRUD request to user repository
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url https://github.com/mihhha985/proxydb
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api
func main() {
	var config = config.LoadConfig()
	var db = db.NewDB(config)
	//repositories
	userRepo := user.NewUserRepository(db)
	//controllers
	userController := user.NewUserController(userRepo)

	// Initialize router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Route("/api", func(r chi.Router) {
		// User routes
		r.Get("/users", userController.GetAll)
		r.Get("/users/{id}", userController.GetOne)
		r.Post("/users", userController.Create)
		r.Put("/users/{id}", userController.Update)
		r.Delete("/users/{id}", userController.Delete)
	})

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
