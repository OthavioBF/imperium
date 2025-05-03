package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	// csrfMiddleware := csrf.Protect(
	// 	[]byte(os.Getenv("GOBID_CSRF_KEY")),
	// 	csrf.Secure(false), // DEV ONLY
	// )
	//
	// api.Router.Use(csrfMiddleware)

	api.Router.Route("/api", func(r chi.Router) {
		// r.Get("/csrftoken", api.HandleGetCSRFtoken)
		r.Route("/users", func(r chi.Router) {
			r.Get("/", api.handleGetUsers)
			r.Get("/{id}", api.handleGetUserById)
			r.Post("/create", api.handleCreateUser)
			r.Put("/{id}", api.handleUpdateUser)
			r.Delete("/{id}", api.handleDeleteUser)
		})
	})
}
