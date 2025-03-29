package repositories

import (
	"context"
	"github.com/dpnam2112/go-backend-template/internal/database"
	"github.com/dpnam2112/go-backend-template/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserRepository handles database operations
type UserRepository struct {
	queries *database.Queries
}

// NewUserRepository initializes a new repository with pgxpool
func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		queries: database.New(pool),
	}
}

// GetUserByID retrieves a user by ID
func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := r.queries.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

// GetUserByID retrieves a user by ID
func (r *UserRepository) CreateUser(ctx context.Context, name string) (*models.User, error) {
	user, err := r.queries.CreateUser(ctx, name)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *UserRepository) WithUnitOfWork(uow *UnitOfWork) *UserRepository {
	userRepo := UserRepository{r.queries.WithTx(uow.tx)}
	return &userRepo
}
