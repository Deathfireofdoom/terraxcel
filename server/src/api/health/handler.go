// Package health contains the health check handlers.
package health

import (
	"fmt"
	"net/http"
)

// HealthCheckHandler responds to health check requests.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type for the response
	w.Header().Set("Content-Type", "text/plain")

	// Write HTTP status code
	w.WriteHeader(http.StatusOK)

	// logging
	fmt.Println("Health check request received")

	// Write response body
	_, err := w.Write([]byte("Everything is ok!"))
	if err != nil {
		// Log the error, replace with your logger of choice
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
