package sheet

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Deathfireofdoom/excel-client-go/pkg/client"
	"github.com/Deathfireofdoom/excel-client-go/pkg/models"

	"github.com/go-chi/chi"
)

// CreateSheetHandler handles the creation of a sheet within a workbook.
func CreateSheetHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the workbook ID from the URL
	workbookID := chi.URLParam(r, "workbookID")

	// Get the sheet from the request body
	var sheet models.Sheet
	err := json.NewDecoder(r.Body).Decode(&sheet)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required sheet fields
	if sheet.Name == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Create the sheet using ExcelClient
	newSheet, err := excelClient.CreateSheet(workbookID, sheet.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create sheet: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert the sheet to JSON
	response, err := json.Marshal(newSheet)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// return the sheet in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetSheetHandler handles fetching a sheet by its ID within a workbook.
func GetSheetHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the workbook ID and sheet ID from the URL
	workbookID := chi.URLParam(r, "workbookID")
	sheetID := chi.URLParam(r, "sheetID")

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Get the sheet using ExcelClient
	sheet, err := excelClient.ReadSheet(workbookID, sheetID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get sheet: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert the sheet to JSON
	response, err := json.Marshal(sheet)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// return the sheet in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// UpdateSheetHandler handles the updating of a sheet by its ID within a workbook.
func UpdateSheetHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the workbook ID and sheet ID from the URL
	workbookID := chi.URLParam(r, "workbookID")
	sheetID := chi.URLParam(r, "sheetID")

	// Get the sheet from the request body
	var sheet models.Sheet
	err := json.NewDecoder(r.Body).Decode(&sheet)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required sheet fields
	if sheet.Name == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Check if the sheet ID in the request body matches the sheet ID in the URL
	if sheet.ID != sheetID {
		http.Error(w, "Sheet ID in request body does not match sheet ID in URL", http.StatusBadRequest)
		return
	}

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Update the sheet using ExcelClient
	updatedSheet, err := excelClient.UpdateSheet(workbookID, &sheet)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update sheet: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert the sheet to JSON
	response, err := json.Marshal(updatedSheet)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// return the sheet in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// DeleteSheetHandler handles the deletion of a sheet by its ID within a workbook.
func DeleteSheetHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the workbook ID and sheet ID from the URL
	workbookID := chi.URLParam(r, "workbookID")
	sheetID := chi.URLParam(r, "sheetID")

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Delete the sheet using ExcelClient
	err = excelClient.DeleteSheet(workbookID, sheetID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete sheet: %v", err), http.StatusInternalServerError)
		return
	}

	// return the sheet in the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted sheet with ID: %s in workbook with ID: %s", sheetID, workbookID)))
}
