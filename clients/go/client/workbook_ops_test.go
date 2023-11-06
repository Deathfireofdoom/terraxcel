package client_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Deathfireofdoom/terraxcel/clients/go/client"
	"github.com/Deathfireofdoom/terraxcel/common/models"
	"github.com/stretchr/testify/require"
)

func TestCreateWorkbook(t *testing.T) {
	// Create a sample workbook to use in the test
	testWorkbook := &models.Workbook{
		ID: "123",
		// ... other fields ...
	}

	// Create a test server that mocks your API's behavior
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/workbook", r.URL.Path)
		require.Equal(t, http.MethodPost, r.Method)

		// Decode the request body
		decoder := json.NewDecoder(r.Body)
		reqWorkbook := &models.Workbook{}
		err := decoder.Decode(reqWorkbook)
		require.NoError(t, err)
		require.Equal(t, testWorkbook, reqWorkbook)

		// Write a response
		w.WriteHeader(http.StatusCreated)
		encoder := json.NewEncoder(w)
		err = encoder.Encode(testWorkbook)
		require.NoError(t, err)
	}))
	defer testServer.Close()

	// Create a client with the URL of the test server
	c, err := client.NewClient(&client.ClientConfig{
		BaseURL:   testServer.URL,
		AuthToken: "test-token",
	})
	require.NoError(t, err)

	// Call the function and check the result
	resultWorkbook, err := c.CreateWorkbook(testWorkbook)
	require.NoError(t, err)
	require.Equal(t, testWorkbook, resultWorkbook)
}
