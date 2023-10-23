package models

import (
	"fmt"

	"github.com/Deathfireofdoom/terraxcel/common/utils"
)

// Sheet represents a single worksheet within a spreadsheet.
// It contains a collection of cells, a name, and a position.
type Sheet struct {
	ID         string `json:"id"`
	WorkbookID string `json:"workbook_id"`
	Pos        int    `json:"pos"`   // Position of the sheet within the spreadsheet
	Name       string `json:"name"`  // Name of the sheet
	Cells      []Cell `json:"cells"` // Collection of cells within the sheet
}

// NewSheet creates a new Sheet instance.
// If id is empty, a new UUID will be generated.
func NewSheet(workbookdId string, pos int, name, id string) (*Sheet, error) {
	if id == "" {
		var err error
		id, err = utils.GenerateUUID()
		if err != nil {
			return nil, fmt.Errorf("failed to generate UUID for Sheet: %w", err)
		}
	}

	return &Sheet{
		ID:         id,
		WorkbookID: workbookdId,
		Pos:        pos,
		Name:       name,
		Cells:      nil,
	}, nil
}
