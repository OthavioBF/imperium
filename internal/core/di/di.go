package di

import (
	"log/slog"
	"os"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/othavioBF/imperium/internal/api"
	"github.com/othavioBF/imperium/internal/services"
)

func InjectDependencies(postgresConn *pgxpool.Pool, sessionManager *scs.SessionManager) api.Api {
	logger := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
			},
		),
	)

	userService := services.NewUserService(postgresConn, logger)
	productService := product.NewService(postgres.NewProductRepository(postgresConn), logger)

	return api.Api{
		Router:         chi.NewMux(),
		Session:        sessionManager,
		Logger:         logger,
		UserService:    userService,
		ProductService: productService,
	}
}
