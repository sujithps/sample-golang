package testutil

import (
	"bytes"
	"encoding/json"
	"git.thoughtworks.net/mahadeva/sample-golang/internal/dependency"
	"git.thoughtworks.net/mahadeva/sample-golang/internal/router"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/instrumentation/mocks"
	"github.com/newrelic/go-agent"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
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
