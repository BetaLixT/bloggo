package mw

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/betalixt/bloggo/svc"
	"github.com/betalixt/bloggo/util/blerr"
)

func AuthMiddleware(tknSvc *svc.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		authSplit := strings.Split(authHeader, " ")
		if 
			len(authSplit) != 2 || authSplit[0] != "Bearer" {
			ctx.Error(blerr.NewError(blerr.TokenInvalidCode, 401, "")) 
		}
		tknSvc.ValidateToken(authSplit[1])
		ctx.Next()
	}
}
