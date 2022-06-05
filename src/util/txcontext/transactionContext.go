package txcontext

import (
	"github.com/betalixt/bloggo/util/http"
	"go.uber.org/zap"
)

type TransactionContext struct {
  logger *zap.Logger
  httpClient *http.HttpClient
}
