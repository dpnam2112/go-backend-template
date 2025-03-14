package providers

import (
	"github.com/dpnam2112/go-backend-template/internal/config"
	"log"

	"go.uber.org/fx"
)

// ProvideConfig loads the application configuration
func ProvideConfig() *config.Config {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}

// Module exports the config provider
var ConfigModule = fx.Module("config",
	fx.Provide(ProvideConfig),
)
