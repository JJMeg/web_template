package runmode

import "github.com/gin-gonic/gin"

type RunMode string

const (
	Production  = RunMode("production")
	Development = RunMode("development")
	Test        = RunMode("test")
)

func (mode RunMode) IsValid() bool {
	switch mode {
	case Production, Development, Test:
		return true
	default:
		return false
	}
}

func (mode RunMode) ParseGinMode() string {
	switch mode {
	case Production:
		return gin.ReleaseMode
	case Development:
		return gin.DebugMode
	case Test:
		return gin.TestMode
	default:
		return ""
	}
}
