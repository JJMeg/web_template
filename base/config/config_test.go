package config

import (
	"github.com/JJMeg/web_template/base/appconfig"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	assertion := assert.New(t)

	var cfg *appconfig.AppConfig

	err := Load("test", "..", &cfg)
	assertion.Nil(err)
	assertion.Nil(cfg.Server)
	assertion.Nil(cfg.Logger)

	err = Load("", "..", &cfg)
	assertion.Nil(err)
	assertion.NotNil(cfg.Server)
	assertion.NotNil(cfg.Logger)

}
