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

func TestCreateSheet(t *testing.T) {
	// Define a sample Sheet object
	testSheet := &models.Sheet{
		ID:         "123",
		WorkbookID: "456",
		// ... other fields ...
	}

	// Mock HTTP server
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/workbook/456/sheet", r.URL.Path)
		require.Equal(t, http.MethodPost, r.Method)

		// Decode request
		decoder := json.NewDecoder(r.Body)
		reqSheet := &models.Sheet{}
		err := decoder.Decode(reqSheet)
		require.NoError(t, err)
		require.Equal(t, testSheet, reqSheet)

		// Write response
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		err = encoder.Encode(testSheet)
		require.NoError(t, err)
	}))
	defer testServer.Close()

	// Client configuration
	c, err := client.NewClient(&client.ClientConfig{
		BaseURL:   testServer.URL,
		AuthToken: "test-token",
	})
	require.NoError(t, err)

	// Call CreateSheet and validate
	resultSheet, err := c.CreateSheet(testSheet)
	require.NoError(t, err)
	require.Equal(t, testSheet, resultSheet)
}
