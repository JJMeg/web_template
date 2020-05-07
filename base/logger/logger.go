package logger

import (
	"github.com/sirupsen/logrus"
)

var ReqIdName = "ReqID"

func NewLogger(cfg *Config) (*logrus.Logger, error) {
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.Level = level

	// output
	out, err := cfg.GetOutputWriter()
	if err != nil {
		return nil, err
	}
	logger.SetOutput(out)

	// formatter
	logger.SetFormatter(cfg.GetFormatter())

	return logger, nil
}

func NewAppLogger(logger *logrus.Logger, reqId string) *logrus.Entry {
	// Fields is map[string]interface{}
	return logger.WithFields(logrus.Fields{
		ReqIdName: reqId,
	})
}

func GetReqId(entry *logrus.Entry) string {
	id, ok := entry.Data[ReqIdName].(string)
	if ok {
		return id
	}

	return ""
}
