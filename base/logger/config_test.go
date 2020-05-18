package logger

import (
	"github.com/sirupsen/logrus"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assertion := assert.New(t)

	cfg := Config{
		Output: "./mock/log",
	}
	assertion.False(cfg.IsStdout())
	out, err := cfg.GetOutputWriter()
	assertion.NotNil(err)
	assertion.Nil(out)

	cfg.Output = "stdout"
	assertion.True(cfg.IsStdout())
	out, err = cfg.GetOutputWriter()
	assertion.Nil(err)
	assertion.NotNil(out)

	cfg.Format = "text"
	assertion.False(cfg.IsJsonFormat())

	cfg.Format = "json"
	assertion.True(cfg.IsJsonFormat())

	formatter := cfg.GetFormatter()
	assertion.NotNil(formatter)
	_, ok := formatter.(*logrus.JSONFormatter)
	assertion.True(ok)
}
