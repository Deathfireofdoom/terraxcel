package workbook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/Deathfireofdoom/excel-client-go/pkg/client"
	"github.com/Deathfireofdoom/excel-client-go/pkg/models"
)

// CreateWorkbookHandler handles the creation of a workbook.
func CreateWorkbookHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Handle if the workbook already exists.
	// Decode the incoming workbook JSON
	var workbook models.Workbook
	err := json.NewDecoder(r.Body).Decode(&workbook)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required workbook fields
	if workbook.FolderPath == "" || workbook.FileName == "" || workbook.Extension == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Create the workbook using ExcelClient
	newWorkbook, err := excelClient.CreateWorkbook(
		workbook.FolderPath,
		workbook.FileName,
		string(workbook.Extension),
		workbook.ID,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create workbook: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the created workbook state in response
	response, err := json.Marshal(newWorkbook)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	fmt.Println("Created workbook: ", newWorkbook.ID)
	fmt.Println("At folder path: ", newWorkbook.FolderPath)
	fmt.Println("With file name: ", newWorkbook.FileName)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetWorkbookHandler handles fetching a workbook by ID.
func GetWorkbookHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL parameters
	id := chi.URLParam(r, "id")

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Fetch the workbook using ExcelClient
	workbook, err := excelClient.ReadWorkbook(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch workbook: %v", err), http.StatusNotFound)
		return
	}

	// Convert the workbook object to JSON
	response, err := json.Marshal(workbook)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// UpdateWorkbookHandler handles the updating of a workbook by ID.
func UpdateWorkbookHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL parameters
	id := chi.URLParam(r, "id")

	// Gets the workbook from the payload
	var workbook models.Workbook
	err := json.NewDecoder(r.Body).Decode(&workbook)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required workbook fields
	if workbook.FolderPath == "" || workbook.FileName == "" || workbook.Extension == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Check if id is the same as the one in the payload
	if id != workbook.ID {
		http.Error(w, "ID in payload does not match ID in URL", http.StatusBadRequest)
		return
	}

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Update the workbook using ExcelClient
	updatedWorkbook, err := excelClient.UpdateWorkbook(&workbook)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update workbook: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the updated workbook state in response
	response, err := json.Marshal(updatedWorkbook)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// DeleteWorkbookHandler handles the deletion of a workbook by ID.
func DeleteWorkbookHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL parameters
	id := chi.URLParam(r, "id")

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Delete the workbook using ExcelClient
	err = excelClient.DeleteWorkbook(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete workbook: %v", err), http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted workbook with ID: %s", id)))
}
