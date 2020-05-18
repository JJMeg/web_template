package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	assertion := assert.New(t)

	cfg := &Config{
		Level:  "info",
		Output: "stdout",
	}

	logger, err := NewLogger(cfg)
	assertion.Nil(err)
	assertion.NotNil(logger)

	cfg.Level = "mock"
	logger, err = NewLogger(cfg)
	assertion.NotNil(err)
	assertion.Nil(logger)

	cfg.Output = "./mock/log"
	logger, err = NewLogger(cfg)
	assertion.NotNil(err)
	assertion.Nil(logger)
}

func TestNewAppLogger(t *testing.T) {
	assertion := assert.New(t)

	cfg := &Config{
		Level:  "info",
		Output: "stdout",
	}
	logger, err := NewLogger(cfg)
	assertion.Nil(err)

	appLogger := NewAppLogger(logger, "mockReqId")
	assertion.NotNil(appLogger)
	assertion.Equal(appLogger.Data["ReqID"],"mockReqId")
}
