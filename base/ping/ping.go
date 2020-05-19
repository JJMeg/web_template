package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []byte("hello!"))
}
