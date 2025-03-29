package providers

import (
	"github.com/dpnam2112/go-backend-template/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

// ProvideUserRepository initializes the user repository
func ProvideUserRepository(dbPool *pgxpool.Pool) *repositories.UserRepository {
	return repositories.NewUserRepository(dbPool)
}

func ProvideUnitOfWorkFactory(dbPool *pgxpool.Pool) *repositories.UnitOfWorkFactory {
	uowFactory := repositories.NewUnitOfWorkFactory(dbPool)
	return uowFactory
}

// RepositoryModule provides all repositories
var RepositoriesModule = fx.Module("repositories",
	fx.Provide(
		ProvideUserRepository,
		ProvideUnitOfWorkFactory,
	),
)
