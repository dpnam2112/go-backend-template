package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds application configuration
type Config struct {
	Host        string `mapstructure:"HOST"`
	Port        int    `mapstructure:"PORT"`
	PostgresURI string `mapstructure:"POSTGRES_URI"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
}

// LoadConfig reads and parses the configuration
func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.SetConfigType("env")    // Read from a .env-like file
	viper.AddConfigPath(path)     // Look for the file in the specified directory
	viper.AutomaticEnv()          // Override with environment variables

	// Default values (optional)
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("LOG_LEVEL", "info")

	// Read config file if available
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found, using environment variables")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
