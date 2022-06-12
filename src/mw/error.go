package mw

import (
	"errors"

	"github.com/betalixt/bloggo/util/blerr"
	"github.com/betalixt/bloggo/util/txcontext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    tctx := ctx.MustGet("tctx").(*txcontext.TransactionContext)
    ctx.Next()
    lgr := tctx.GetLogger()    
    if len(ctx.Errors) > 0 {
      errs := make([]error, len(ctx.Errors))
      berr := (*blerr.Error)(nil)
      for idx, err := range ctx.Errors {
        errs[idx] = err.Err
        var temp *blerr.Error
        if errors.As(err.Err, &temp) {
          berr = temp
        }
      }
      lgr.Error("errors processing request", zap.Errors("error", errs))
      if berr != nil {
        ctx.JSON(berr.StatusCode, berr)
      } else {
        ctx.JSON(500, blerr.UnexpectedError())
      }
    } else {
      if (!ctx.Writer.Written()) {
        lgr.Error("No response was written")
        ctx.JSON(500, blerr.UnsetResponseError())
      }
    }

  }
}
