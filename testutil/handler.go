package testutil

import (
	"bytes"
	"encoding/json"
	"github.com/newrelic/go-agent"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"spikes/sample-golang/internal/dependency"
	"spikes/sample-golang/internal/router"
	"spikes/sample-golang/pkg/instrumentation/mocks"
	"testing"
)

func InitTestContainer() *dependency.Container {
	return dependency.Init(mockNewRelicApp)
}

func StartTestServer(container *dependency.Container) *httptest.Server {
	rtr := router.Router(container)
	return httptest.NewServer(rtr)
}

func mockNewRelicApp() newrelic.Application {
	return &mocks.StubNewrelicApp{}
}

func MakeHTTPCall(method string, url string, requestBody interface{}) (*http.Response, error) {
	var body io.Reader
	if requestBody != nil {
		byt, _ := json.Marshal(requestBody)
		body = bytes.NewBuffer(byt)
	} else {
		body = nil
	}

	request, _ := http.NewRequest(method, url, body)
	client := http.Client{}
	return client.Do(request)
}

func ParseResponseBody(t *testing.T, responseBody io.ReadCloser, apiResponse interface{}) {
	err := json.NewDecoder(responseBody).Decode(apiResponse)
	assert.NoError(t, err)
}
