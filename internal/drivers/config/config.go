package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config represents config for Easypark app.
type Config struct {
	Database  DatabaseConfig
	HTTP      HTTPConfig
	Auth      AuthConfig
	Logging   LoggingConfig
	Scheduler SchedulerConfig
	Alert     AlertConfig
}

// DatabaseConfig represents a config for our Database.
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// HTTPConfig represents a config for our HTTP server.
type HTTPConfig struct {
	Address          string
	WebsocketAddress string
}

// AuthConfig represents our config for Auth service.
type AuthConfig struct {
	SecretKey string
}

// LoggingConfig represents a config for our Logger.
type LoggingConfig struct {
	Level     string
	GormLevel string
	Format    string
}

// SchedulerConfig represents a config for our Scheduler.
type SchedulerConfig struct {
	Interval string
}

// AlertConfig represents a config for alerts.
type AlertConfig struct {
	LateArrivalThresholdMinutes time.Duration
	OverStayThresholdMinutes    time.Duration
}

// LoadConfig reads configuration from ./config.yaml file.
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
