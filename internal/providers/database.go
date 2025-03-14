package providers

import (
	"context"
	"github.com/dpnam2112/go-backend-template/internal/config"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"go.uber.org/fx"
)

// ProvidePgConnPool provides database connection pool
func ProvidePgConnPool(cfg *config.Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cfg.PostgresURI)
	if err != nil {
		return nil, err
	}

	// Set up logger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger: tracelog.LoggerFunc(func(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
			logger.Printf("[%s] %s: %v", level, msg, data)
		}),
		LogLevel: tracelog.LogLevelDebug, // Adjust log level as needed
	}

	// Create connection pool
	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}

// Module exports the database provider
var DatabaseModule = fx.Module("database",
	fx.Provide(ProvidePgConnPool),
)
