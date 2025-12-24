package config

// Config holds application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
	Host string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// TODO: Implement configuration loading
	return &Config{}, nil
}
