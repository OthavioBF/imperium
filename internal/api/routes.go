package api

import (
	"net/http"

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
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
		r.Route("/users", func(r chi.Router) {
			r.Get("/", api.handleGetUsers)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", api.handleGetUserById)
				r.Post("/create", api.handleCreateUser)
				r.Put("/update", api.handleUpdateUser)
				r.Delete("/delete", api.handleDeleteUser)
			})
		})
	})
}
