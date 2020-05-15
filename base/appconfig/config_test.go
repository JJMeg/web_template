package appconfig

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var appCfg = `
{
    "name":"mock_app",
    "server":{
        "host":"0.0.0.0:3000",
        "request_timeout":0,
        "response_timeout":0,
        "request_max":50000,
        "throttle":10000,
        "request_id":"xx",
        "request_pin":"xx"
    },
    "logger":{
        "output":"stdout",
        "level":"debug"
    }
}
`

func TestAppConfig(t *testing.T) {
	assertion := assert.New(t)

	var cfg *AppConfig

	err := json.Unmarshal([]byte(appCfg), &cfg)
	assertion.Nil(err)

	copyCfg := cfg.Copy()
	assertion.Equal(fmt.Sprintf("%+v", cfg), fmt.Sprintf("%+v", copyCfg))
	assertion.NotEqual(fmt.Sprintf("%p", cfg), fmt.Sprintf("%p", copyCfg))

	assertion.Equal(cfg.Name, cfg.GetAppName())
	assertion.Equal(cfg.Server.Throttle, cfg.GetThrottle())
	assertion.Equal(cfg.Server.RequestMax, cfg.GetRequestMax())
}
