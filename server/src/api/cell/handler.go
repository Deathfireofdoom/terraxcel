// api/cell/cell.go
package cell

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Deathfireofdoom/terraxcel/common/models"
	client "github.com/Deathfireofdoom/terraxcel/server/src/pkg/terraxcel"

	"github.com/go-chi/chi"
)

// CreateCellHandler handles the creation of a cell.
func CreateCellHandler(w http.ResponseWriter, r *http.Request) {
	// Get the workbookID and sheetID from the URL
	workbookID := chi.URLParam(r, "workbookID")
	sheetID := chi.URLParam(r, "sheetID")

	// Get the cell from the request body
	var cell models.Cell
	err := json.NewDecoder(r.Body).Decode(&cell)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required cell fields TODO check this
	if cell.Row == 0 || cell.Column == "" || cell.Value == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Initialize the Excel client
	excelClient, err := client.NewTerraxcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Create the cell using ExcelClient
	newCell, err := excelClient.CreateCell(workbookID, sheetID, cell.Row, cell.Column, cell.Value)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create cell: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert the cell to JSON
	response, err := json.Marshal(newCell)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// return the cell in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetCellHandler handles fetching a cell by its ID.
func GetCellHandler(w http.ResponseWriter, r *http.Request) {
	// Get the workbookID, sheetID, and cellID from the URL
	workbookID := chi.URLParam(r, "workbookID")
	sheetID := chi.URLParam(r, "sheetID")
	cellID := chi.URLParam(r, "cellID")

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Get the cell using ExcelClient
	cell, err := excelClient.ReadCell(workbookID, sheetID, cellID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read cell: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert the cell to JSON
	response, err := json.Marshal(cell)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// return the cell in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// UpdateCellHandler handles updating a cell by its ID.
func UpdateCellHandler(w http.ResponseWriter, r *http.Request) {
	// Get the workbookID, sheetID, and cellID from the URL
	workbookID := chi.URLParam(r, "workbookID")
	sheetID := chi.URLParam(r, "sheetID")
	cellID := chi.URLParam(r, "cellID")

	// Get the cell from the request body
	var cell models.Cell
	err := json.NewDecoder(r.Body).Decode(&cell)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required cell fields TODO check this
	if cell.Row == 0 || cell.Column == "" || cell.Value == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Check if the cell ID in the request body matches the cell ID in the URL
	if cell.ID != cellID {
		http.Error(w, "Cell ID in request body does not match cell ID in URL", http.StatusBadRequest)
		return
	}

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Update the cell using ExcelClient
	updatedCell, err := excelClient.UpdateCell(workbookID, sheetID, &cell)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update cell: %v", err), http.StatusInternalServerError)
		return
	}

	// Convert the cell to JSON
	response, err := json.Marshal(updatedCell)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// return the cell in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// DeleteCellHandler handles deleting a cell by its ID.
func DeleteCellHandler(w http.ResponseWriter, r *http.Request) {
	// Get the workbookID, sheetID, and cellID from the URL
	workbookID := chi.URLParam(r, "workbookID")
	sheetID := chi.URLParam(r, "sheetID")
	cellID := chi.URLParam(r, "cellID")

	// Initialize the Excel client
	excelClient, err := client.NewExcelClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to initialize Excel client: %v", err), http.StatusInternalServerError)
		return
	}

	// Delete the cell using ExcelClient
	err = excelClient.DeleteCell(workbookID, sheetID, cellID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete cell: %v", err), http.StatusInternalServerError)
		return
	}

	// return the cell in the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted cell with ID: %s in workbook %s and sheet %s", cellID, workbookID, sheetID)))
}
