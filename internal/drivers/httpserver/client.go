package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	configPath = "." // located in project root.
	configName = "config"
	configType = "yaml"
	configKey  = "httpserver"
)

type ClientConfig struct {
	Port int
	Host string
}

// Client represents an HTTP server client
type Client struct {
	router *gin.Engine
	config ClientConfig
}

func NewClientFromConfig() (*Client, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	var config ClientConfig

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.UnmarshalKey(configKey, &config)
	if err != nil {
		return nil, err
	}

	return newClient(config), nil
}

func newClient(config ClientConfig) *Client {
	router := gin.Default()

	client := &Client{
		router: router,
		config: config,
	}

	// TODO
	// Setup routes

	return client
}

// Run starts the HTTP server
func (c *Client) Run() error {
	addr := fmt.Sprintf("%s:%d", c.config.Host, c.config.Port)
	return c.router.Run(addr)
}
