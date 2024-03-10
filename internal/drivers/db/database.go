package drivers

import (
	"github.com/spf13/viper"
)

const (
	configPath = "." // located in project root.
	configName = "config"
	configType = "yaml"
	configKey  = "database"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// LoadDatabseConfig loads our databse configuration from config.yaml in the project root using Viper.
func LoadDatabaseConfig() (*DatabaseConfig, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	var config DatabaseConfig

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.UnmarshalKey(configKey, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
