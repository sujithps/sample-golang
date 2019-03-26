package profiling

import (
	"github.com/newrelic/go-agent"
	"context"
	"spikes/sample-golang/pkg/appcontext"
)

type Ender interface {
	End() error
}

type noopEnder struct{}

func (*noopEnder) End() error { return nil }

var noop = noopEnder{}

type tracer struct {
	product newrelic.DatastoreProduct
}

var MongoTracer = tracer{product: newrelic.DatastoreMongoDB}
var ESTracer = tracer{product: newrelic.DatastoreElasticsearch}

func (t tracer) Start(ctx context.Context, collection, name string) Ender {
	if txn := appcontext.GetNewRelicTransaction(ctx); txn != nil {
		ds := &newrelic.DatastoreSegment{
			Product:    t.product,
			Operation:  name,
			Collection: collection,
		}
		ds.StartTime = newrelic.StartSegmentNow(txn)
		return ds
	}
	return &noop
}
