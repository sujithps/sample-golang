package handler_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"spikes/sample-golang/internal/dependency"
	"spikes/sample-golang/internal/domain"
	"spikes/sample-golang/pkg/contract"
	"spikes/sample-golang/testutil"
	"testing"
)

var testContainer *dependency.Container

func TestMain(m *testing.M) {
	testutil.Setup()
	testContainer = testutil.InitTestContainer()
	code := m.Run()
	os.Exit(code)
}

func TestShouldReturn4XXErrorForUserNotFoundWhenInternalGetUserCalled(t *testing.T) {
	testServer := testutil.StartTestServer(testContainer)
	defer testServer.Close()

	response, err := testutil.MakeHTTPCall("GET", fmt.Sprintf("%s/user/hal",
		testServer.URL), nil)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	res := &contract.Response{}
	testutil.ParseResponseBody(t, response.Body, res)
	assert.Nil(t, res.Data)
}

func TestShould_ReturnUserFromMongo(t *testing.T) {
	user := domain.NewUser("123", "first", "last")
	defer testutil.CleanDb()

	_ = testContainer.GetUserService().Upsert(context.Background(), user)

	testServer := testutil.StartTestServer(testContainer)
	defer testServer.Close()

	response, err := testutil.MakeHTTPCall("GET", fmt.Sprintf("%s/user/123", testServer.URL), nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	res := &APIResponse{}
	testutil.ParseResponseBody(t, response.Body, res)
	assert.True(t, res.Success)
	assert.Equal(t, user.ID, res.Data.ID)
	assert.Equal(t, "last, first", res.Data.DisplayName)
}

type APIResponse struct {
	Data    contract.User    `json:"data"`
	Success bool             `json:"success"`
	Errors  []contract.Error `json:"errors,omitempty"`
}
