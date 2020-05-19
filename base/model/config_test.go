package model

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assertion := assert.New(t)

	appCfg := `
	{
		"host":"localhost:27017",
		"user":"root",
		"passwd":"passwd",
		"database":"test",
		"mode":"strong",
		"pool":5,
		"timeout":5,
		"replica":"mgset-500148149"
	}
	`

	modelCfg, err := NewConfig([]byte(appCfg))
	assertion.Nil(err)

	assertion.Equal("test", modelCfg.Database)
	assertion.Equal("root", modelCfg.GetUser())
	assertion.Equal("passwd", modelCfg.GetPasswd())

	// inject env params
	modelCfg.User = ""
	err = os.Setenv("MONGODB_USER", "MONGODB_USER")
	assertion.Nil(err)
	assertion.Equal("MONGODB_USER", modelCfg.GetUser())

	modelCfg.Passwd = ""
	err = os.Setenv("MONGODB_PASSWD", "MONGODB_PASSWD")
	assertion.Nil(err)
	assertion.Equal(modelCfg.GetPasswd(), "MONGODB_PASSWD")

	// test copy()
	cfg := new(Config)
	copiedCfg := cfg.Copy()

	assertion.NotEqual(fmt.Sprintf("%p", cfg), fmt.Sprintf("%p", copiedCfg))
}
