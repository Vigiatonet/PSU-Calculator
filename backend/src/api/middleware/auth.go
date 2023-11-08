package middleware

import (
	"net/http"
	"strings"

	"github.com/Vigiatonet/PSU-Calculator/api/helper"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/constants"
	"github.com/Vigiatonet/PSU-Calculator/pkg/service_errors"
	"github.com/Vigiatonet/PSU-Calculator/services"
	"github.com/gin-gonic/gin"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenService = services.NewTokenService(cfg)
	var err error
	var claimMap = map[string]interface{}{}
	return func(ctx *gin.Context) {
		tk := ctx.GetHeader(constants.AuthenTicationHeaderKey)
		if tk == "" {
			err = &service_errors.ServiceError{EndUserMsg: service_errors.TokenNotPresent}

		} else {
			token := strings.Split(tk, " ")[1]
			claimMap, err = tokenService.GetClaims(token)
			if err != nil {
				e, isServiceError := err.(*service_errors.ServiceError)
				if isServiceError && e.EndUserMsg == service_errors.TokenExpired {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError("provided Token is expired", helper.AuthError, false, err))
					return
				} else {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(nil, helper.AuthError, false, err))
					return
				}
			}
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(nil, -1, false, err))
			return
		}
		ctx.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		ctx.Set(constants.PhoneKey, claimMap[constants.PhoneKey])
		ctx.Set(constants.ExpKey, claimMap[constants.ExpKey])
		ctx.Set(constants.RolesKey, claimMap[constants.RolesKey])

		ctx.Next()
	}
}
