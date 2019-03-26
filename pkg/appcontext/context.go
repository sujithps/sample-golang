package appcontext

import (
	"context"
	"github.com/newrelic/go-agent"
	"net/http"
)

type Key string

const (
	contextNewRelicKey = Key("NewRelicTxn")
	correlationIDKey   = Key("CorrelationID")
)

func WithNewRelicTransaction(r *http.Request, transaction newrelic.Transaction) *http.Request {
	c := context.WithValue(r.Context(), contextNewRelicKey, transaction)
	return r.WithContext(c)
}

func WithCorrelationID(r *http.Request, correlationID string) *http.Request {
	c := context.WithValue(r.Context(), correlationIDKey, correlationID)
	return r.WithContext(c)
}

func NewBackgroundContext(transaction newrelic.Transaction, correlationID string) context.Context {
	ctx := context.Background()
	c := context.WithValue(ctx, contextNewRelicKey, transaction)
	c = context.WithValue(c, correlationIDKey, correlationID)
	return c
}

func GetNewRelicTransaction(ctx context.Context) newrelic.Transaction {
	if val, ok := ctx.Value(contextNewRelicKey).(newrelic.Transaction); ok {
		return val
	}
	return nil
}

func GetCorrelationID(ctx context.Context) string {
	if val, ok := ctx.Value(correlationIDKey).(string); ok {
		return val
	}
	return ""
}
