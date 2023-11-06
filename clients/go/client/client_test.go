package client_test

import (
	"testing"

	"github.com/Deathfireofdoom/terraxcel/clients/go/client"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		config  *client.ClientConfig
		wantErr bool
	}{
		{
			name: "missing AuthToken",
			config: &client.ClientConfig{
				BaseURL: "http://example.com",
			},
			wantErr: true,
		},
		{
			name: "missing BaseURL",
			config: &client.ClientConfig{
				AuthToken: "test-token",
			},
			wantErr: true,
		},
		{
			name: "valid config",
			config: &client.ClientConfig{
				BaseURL:   "http://example.com",
				AuthToken: "test-token",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.NewClient(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewRequest(t *testing.T) {
	validClient, err := client.NewClient(&client.ClientConfig{
		BaseURL:   "http://example.com",
		AuthToken: "test-token",
	})
	if err != nil {
		t.Fatalf("Failed to create valid client: %v", err)
	}

	tests := []struct {
		name     string
		client   *client.Client
		method   string
		endpoint string
		body     interface{}
		wantErr  bool
	}{
		{
			name:     "valid request",
			client:   validClient,
			method:   "POST",
			endpoint: "/test",
			body: map[string]string{
				"key": "value",
			},
			wantErr: false,
		},
		// Add more test cases as needed for different scenarios
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.client.NewRequest(tt.method, tt.endpoint, tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
