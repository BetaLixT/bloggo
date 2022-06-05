package txcontext

import (
	"github.com/betalixt/bloggo/util/http"
	"go.uber.org/zap"
)

type TransactionContext struct {
	cid        string
	rid        string
	logger     *zap.Logger
	httpClient *http.HttpClient
}

func (tctx *TransactionContext) GetLogger() *zap.Logger {
	return tctx.logger
}
func (tctx *TransactionContext) GetHttpClient() *http.HttpClient {
	if tctx.httpClient == nil {
		tctx.httpClient = http.NewClient(
			tctx.logger,
			map[string]string{
				"x-correlation-id": tctx.cid,
			},
		)
	}
	return tctx.httpClient
}

// - Constructor
func NewTransactionContext(
	cid string,
	rid string,
	logger *zap.Logger,
) *TransactionContext {

	return &TransactionContext{
		cid: cid,
		rid: rid,
		logger: logger.With(zap.String("cid", cid), zap.String("rid", rid)),
	}
}
