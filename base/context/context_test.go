package context

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/JJMeg/web_template/base/logger"
)

func TestContext(t *testing.T) {
	assertion := assert.New(t)

	ctx := &gin.Context{}
	log := NewLogger(ctx)
	assertion.Nil(log)

	log = NewLogger(newMockCtxWithLogger())
	assertion.NotNil(log)
}

func TestContext_Logger(t *testing.T) {
	assertion := assert.New(t)

	ctx := Context{}
	assertion.Nil(ctx.logger)

	mockCtx := newMockCtxWithLogger()
	ctx.Context = mockCtx
	assertion.NotNil(ctx.Logger())
}

func TestNewHandler(t *testing.T) {
	assertion := assert.New(t)

	fn := func(ctx *Context) {}
	handler := NewHandler(fn)

	handler(&gin.Context{})

	assertion.NotNil(handler)
	assertion.Equal(fmt.Sprintf("%T", handler), "gin.HandlerFunc")
}

func TestNewLoggerMiddleware(t *testing.T) {
	assertion := assert.New(t)

	l := newMockLogger()

	handler := NewLoggerMiddleware(l)
	assertion.NotNil(handler)
	assertion.Equal(fmt.Sprintf("%T", handler), "gin.HandlerFunc")

	ctx := &gin.Context{
		Request: &http.Request{
			Header: http.Header{},
		},
	}
	handler(ctx)
	assertion.NotNil(ctx.Get(ContextLoggerKey))
}

func newMockCtxWithLogger() *gin.Context {
	l := newMockLogger()

	ctx := &gin.Context{}
	ctx.Set(ContextLoggerKey, logger.NewAppLogger(l, "reqid"))

	return ctx
}

func newMockLogger() *logrus.Logger {
	l, _ := logger.NewLogger(&logger.Config{
		Output: "stdout",
		Level:  "info",
	})

	return l
}
