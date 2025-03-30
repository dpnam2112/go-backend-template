// @title dpnam2112's Go Backend Template
// @version 1.0
// @description This is a backend template for building scalable and modular Go applications.
//
// @contact.name Support Team
// @contact.url https://github.com/dpnam2112
// @contact.email dpnam2112@gmail.com
//
// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
//
// @BasePath /
// @schemes http https

package main

import (
	"context"
	"fmt"
	"github.com/dpnam2112/go-backend-template/internal/config"
	"github.com/dpnam2112/go-backend-template/internal/providers"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

import _ "github.com/dpnam2112/go-backend-template/docs"

// StartServer runs the Gin HTTP server using config values
func StartServer(lc fx.Lifecycle, router *gin.Engine, cfg *config.Config) {
	// Construct server address from config
	serverAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Starting Gin server on %s", serverAddr)

			// Run the Gin server in a goroutine to prevent blocking
			go func() {
				if err := router.Run(serverAddr); err != nil {
					log.Fatalf("Gin server failed to start: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down Gin server")
			return nil // Gin handles graceful shutdown automatically
		},
	})
}

func main() {
	app := fx.New(
		providers.ConfigModule,
		providers.DatabaseModule,
		providers.RepositoriesModule,
		providers.HandlersModule,
		providers.RouterModule,
		providers.LoggerModule,

		// Start Gin server with config-based host/port
		fx.Invoke(StartServer),
	)

	app.Run()
}
