package middleware

import (
	"net/http"

	"github.com/Vigiatonet/PSU-Calculator/api/helper"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(2, nil)
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, -1, false, err))
			return
		}
		ctx.Next()
	}
}
