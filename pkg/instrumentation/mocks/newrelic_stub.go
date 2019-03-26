package mocks

import (
	"net/http"
	"time"

	"github.com/newrelic/go-agent"
)

type StubNewRelicTransaction struct{
	http.ResponseWriter
}

func (snrt *StubNewRelicTransaction) End() error {
	return nil
}

func (snrt *StubNewRelicTransaction) Ignore() error {
	return nil
}

func (snrt *StubNewRelicTransaction) SetName(name string) error {
	return nil
}

func (snrt *StubNewRelicTransaction) NoticeError(err error) error {
	return nil
}

func (snrt *StubNewRelicTransaction) AddAttribute(key string, value interface{}) error {
	return nil
}

func (snrt *StubNewRelicTransaction) StartSegmentNow() newrelic.SegmentStartTime {
	return newrelic.SegmentStartTime{}
}

func (snrt *StubNewRelicTransaction) Header() http.Header {
	return http.Header{}
}

func (snrt *StubNewRelicTransaction) Write(a []byte) (int, error) {
	return snrt.ResponseWriter.Write(a)
}

func (snrt *StubNewRelicTransaction) WriteHeader(s int) {
	snrt.ResponseWriter.WriteHeader(s)
	return
}

type StubNewrelicApp struct{}

func (sna *StubNewrelicApp) StartTransaction(name string, w http.ResponseWriter, r *http.Request) newrelic.Transaction {
	return &StubNewRelicTransaction{w}
}

func (sna *StubNewrelicApp) RecordCustomEvent(eventType string, params map[string]interface{}) error {
	return nil
}

func (sna *StubNewrelicApp) RecordCustomMetric(name string, value float64) error {
	return nil
}

func (sna *StubNewrelicApp) WaitForConnection(timeout time.Duration) error {
	return nil
}
func (sna *StubNewrelicApp) Shutdown(timeout time.Duration) {
	return
}
