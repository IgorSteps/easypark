package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

// NewDatabaseFromConfig creates our databse from config.
func NewDatabaseFromConfig() (*gorm.DB, error) {
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

	return newDatabase(config)
}

func newDatabase(config DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	return db, nil
}
