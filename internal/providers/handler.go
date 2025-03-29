package providers

import (
	"github.com/dpnam2112/go-backend-template/internal/handlers"
	"github.com/dpnam2112/go-backend-template/internal/repositories"

	"go.uber.org/fx"
)

// ProvideUserHandler initializes the user handler
func ProvideUserHandler(userRepo *repositories.UserRepository, uowFactory *repositories.UnitOfWorkFactory) *handlers.UserHandler {
	return handlers.NewUserHandler(userRepo, uowFactory)
}

// HandlersModule provides all HTTP handlers
var HandlersModule = fx.Module("handlers",
	fx.Provide(
		ProvideUserHandler,
	),
)
