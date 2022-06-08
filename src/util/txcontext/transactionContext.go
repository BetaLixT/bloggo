package txcontext

import (
	"github.com/betalixt/bloggo/intl/db"
	"github.com/betalixt/bloggo/intl/trace"
	"github.com/betalixt/bloggo/intl/http"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type TransactionContext struct {
	cid        string
	rid        string
	isParent   bool
	db         *sqlx.DB
	logger     *zap.Logger
	tracer     trace.ITracer
	httpClient *http.HttpClient
	tracedDB   *db.TracedDBContext
}

func (tctx *TransactionContext) GetLogger() *zap.Logger {
	return tctx.logger
}
func (tctx *TransactionContext) GetHttpClient() *http.HttpClient {
	if tctx.httpClient == nil {
		tctx.httpClient = http.NewClient(
			tctx.GetTracer(),
			map[string]string{
				"x-correlation-id": tctx.cid,
			},
		)
	}
	return tctx.httpClient
}
func (tctx *TransactionContext) GetDatabaseContext() *db.TracedDBContext {
	if tctx.tracedDB == nil {
		tctx.tracedDB = db.NewTracedDBContext(
			tctx.db,
			tctx.GetTracer(),
			"main-database",
		)
	}
	return tctx.tracedDB
}

func (tctx *TransactionContext) GetTracer() trace.ITracer {
	if tctx.tracer == nil {
		tctx.tracer = trace.NewZapTracer(tctx.logger)
	}
	return tctx.tracer
}

func (tctx *TransactionContext) IsParent() bool {
	return tctx.isParent
}

// - Constructor
func NewTransactionContext(
	cid string,
	rid string,
	isParent bool,
	db *sqlx.DB,
	logger *zap.Logger,
) *TransactionContext {

	return &TransactionContext{
		cid:      cid,
		rid:      rid,
		isParent: isParent,
		db:       db,
		logger:   logger.With(zap.String("cid", cid), zap.String("rid", rid)),
	}
}
