package model

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Host       string        `json:"host"`
	User       string        `json:"user"`
	Passwd     string        `json:"passwd"`
	Database   string        `json:"database"`
	Mode       string        `json:"mode"`
	Pool       int           `json:"pool"`
	Timeout    time.Duration `json:"timeout"`
	PEMFILE    string        `json:"pemfile"` //permission file
	ReplicaSet string        `json:"replica"`
}

func NewConfig(data []byte) (*Config, error) {
	var cfg Config
	err := json.Unmarshal(data, &cfg)
	return &cfg, err
}

func (c *Config) GetUser() string {
	if c.User != "" {
		return c.User
	}

	return os.Getenv("MONGODB_USER")
}

func (c *Config) GetPasswd() string {
	if c.Passwd != "" {
		return c.Passwd
	}

	return os.Getenv("MONGODB_PASSWD")
}

func (c *Config) Copy() *Config {
	cfg := *c

	return &cfg
}

func (c *Config) DSN() string {
	dsn := "mongodb://"

	// set user & password
	{
		user := c.GetUser()
		password := c.GetPasswd()

		if user != "" && password != "" {
			dsn += user + ":" + password + "@"
		}
	}

	// set database
	dsn += c.Host
	if c.Database != "" {
		dsn += "/" + c.Database
	}

	// set replica set
	if c.ReplicaSet != "" {
		dsn += fmt.Sprintf("?replicaSet=%s", c.ReplicaSet)
	}

	return dsn
}
