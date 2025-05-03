package api

import (
	"log/slog"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/othavioBF/imperium/internal/services"
)

type Api struct {
	Router      *chi.Mux
	Session     *scs.SessionManager
	Logger      *slog.Logger
	UserService services.UserService
}
