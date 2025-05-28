package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/othavioBF/imperium/internal/infra/pgstore"
)

var (
	ErrCreateUser  = errors.New("Falha ao criar usuario")
	ErrGetUserById = errors.New("Falha ao buscar usuario")
)

type UserService struct {
	pool   *pgxpool.Pool
	repo   pgstore.Querier
	logger *slog.Logger
}

func NewUserService(pool *pgxpool.Pool, logger *slog.Logger) *UserService {
	return &UserService{
		pool:   pool,
		repo:   pgstore.New(pool),
		logger: logger,
	}
}

func (s *UserService) GetUsers(ctx context.Context, id uuid.UUID) ([]pgstore.User, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUserById(ctx context.Context, id uuid.UUID) error {
	// s.logger.Error("User Service", "Critical error")
	return nil
}

func (s *UserService) CreateUser(ctx context.Context, user pgstore.CreateUserParams) (uuid.UUID, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		s.logger.Error("User Service", "Failed to begin transaction: ", err)
		return uuid.UUID{}, err
	}
	defer func() {
		if err != nil {
			if err = tx.Rollback(ctx); err != nil {
				s.logger.Error("User Service", "Rollback error", err)
			}
		}
	}()

	repo := pgstore.New(tx)

	id, err := repo.CreateUser(ctx, pgstore.CreateUserParams{
		UserName: user.UserName,
		Email:    user.Email,
	})
	if err != nil {
		s.logger.Error("User Service", "err", err)
		return uuid.UUID{}, ErrCreateUser
	}

	if err = tx.Commit(ctx); err != nil {
		s.logger.Error("User Service", "Commit error", err)
		return uuid.UUID{}, err
	}

	return id, nil
}
