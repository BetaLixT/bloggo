package mw

import (
	// "strings"

	"github.com/betalixt/bloggo/util/txcontext"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	// "github.com/betalixt/bloggo/util/blerr"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func TransactionContextGenerationMiddleware(
	lgr *zap.Logger,
	db *sqlx.DB,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cid := ctx.GetHeader("x-correlation-id")
		isParent := false
		if cid == "" {
			cid = uuid.NewV4().String()
			isParent = true
		}
		rid := uuid.NewV4().String()
		tctx := txcontext.NewTransactionContext(
			cid,
			rid,
			isParent,
			db,
			lgr,
		)
		ctx.Set("tctx", tctx)
		ctx.Writer.Header().Set("x-correlation-id", cid)
		ctx.Next()
	}
}
