package providers

import (
	"github.com/dpnam2112/go-backend-template/internal/handlers"
	"github.com/dpnam2112/go-backend-template/internal/routers"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// ProvideRouter initializes the Gin router and registers routes
func ProvideRouter(userHandler *handlers.UserHandler) *gin.Engine {
	ginRouter := gin.Default()
	routers.RegisterRoutes(ginRouter, userHandler)
	return ginRouter
}

// RouterModule provides the router dependency
var RouterModule = fx.Module("router",
	fx.Provide(ProvideRouter),
)
