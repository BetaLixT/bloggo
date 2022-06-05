package mw

import (
	"errors"
	"net"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/betalixt/bloggo/util/blerr"
	"go.uber.org/zap"
)

func RecoveryMiddleware(lgr *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				perr, ok := err.(blerr.Error)
				if ok {
					c.JSON(perr.StatusCode, perr)
				} else {
					// Check for a broken connection, as it is not really a
					// condition that warrants a panic stack trace.
					var brokenPipe bool
					if ne, ok := err.(*net.OpError); ok {
						var se *os.SyscallError
						if errors.As(ne, &se) {
							if strings.Contains(
								strings.ToLower(se.Error()), "broken pipe") ||
								strings.Contains(strings.ToLower(se.Error()),
								"connection reset by peer",
								) {
								brokenPipe = true
							}
						}
					}
					
					httpRequest, _ := httputil.DumpRequest(c.Request, false)
					headers := strings.Split(string(httpRequest), "\r\n")
					for idx, header := range headers {
						current := strings.Split(header, ":")
						if current[0] == "Authorization" {
							headers[idx] = current[0] + ": *"
						}
					}
					headersToStr := strings.Join(headers, "\r\n")
					if brokenPipe {	
						lgr.Error(
							"Panic recovered, broken pipe",
							zap.String("headers", headersToStr),
							zap.Any("error", err),
						)
						c.Abort()
					} else {	
						lgr.Error(
							"Panic recovered",
							zap.String("headers", headersToStr),
							zap.Any("error", err),
							zap.Stack("stack"),
						)
						c.JSON(500, blerr.UnexpectedError())
					}
				}
			}
		}()
		c.Next()
	}
}
