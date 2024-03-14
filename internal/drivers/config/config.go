package config

import "github.com/spf13/viper"

type Config struct {
	Database DatabaseConfig
	Auth     AuthConfig
	Logging  LoggingConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type AuthConfig struct {
	SecretKey string
}

type LoggingConfig struct {
	Level     string
	GormLevel string
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
