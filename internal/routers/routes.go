package routers

import (
	"github.com/dpnam2112/go-backend-template/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// PingHander to check server's health
// @Summary Ping the server
// @Description Health check endpoint
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// SetupRoutes sets up the routes for the API
func RegisterRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {

	// Health check endpoint
	router.GET("/ping", PingHandler)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	{
		userRoutes := v1.Group("/users")
		{
			userRoutes.GET("/:id", userHandler.GetUser)
			userRoutes.POST("", userHandler.CreateUser)
		}
	}
}
