package mw

import (
	// "strings"

	"github.com/betalixt/bloggo/util/txcontext"
	"go.uber.org/zap"

	// "github.com/betalixt/bloggo/util/blerr"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func TransactionContextGenerationMiddleware(
	lgr *zap.Logger,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cid := ctx.GetHeader("x-correlation-id")
		if cid == "" {
			cid = uuid.NewV4().String()
		}
		rid := uuid.NewV4().String()
		tctx := txcontext.NewTransactionContext(
			cid,
			rid,
			lgr,
		)
		ctx.Set("tctx", tctx)
		ctx.Writer.Header().Set("x-correlation-id", cid)
		ctx.Next()
	}
}
