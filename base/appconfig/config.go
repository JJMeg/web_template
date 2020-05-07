package appconfig

import (
	"github.com/JJMeg/web_template/base/logger"
)

type ServerConfig struct {
	Host            string `json:"host"`
	RequestTimeout  int    `json:"request_timeout"`
	ResponseTimeout int    `json:"response_timeout"`

	Throttle   int `json:"throttle"`
	RequestMax int `json:"request_max"`

	RequestID  string `json:"request_id"`
	RequestPin string `json:"request_pin"`
}

type AppConfig struct {
	Name   string         `json:"name"`
	Server *ServerConfig  `json:"server"`
	Logger *logger.Config `json:"logger"`
}

func (c *AppConfig) Copy() *AppConfig {
	cfg := *c
	return &cfg
}

func (c *AppConfig) GetAppName() string {
	return c.Name
}

func (c *AppConfig) GetThrottle() int {
	if c.Server.Throttle == 0 {
		return 100
	}

	return c.Server.Throttle
}

func (c *AppConfig) GetRequestMax() int {
	if c.Server.RequestMax == 0 {
		return 25
	}

	return c.Server.RequestMax
}
