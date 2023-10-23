package models

import (
	"fmt"

	"github.com/Deathfireofdoom/terraxcel/common/utils"
)

// Cell represents a single cell in a spreadsheet
type Cell struct {
	ID         string    `json:"id"`
	WorkbookID string    `json:"workbook_id"`
	SheetID    string    `json:"sheet_id"`
	Row        int       `json:"row"`
	Column     string    `json:"column"`
	Value      CellValue `json:"value"`
}

type CellValue struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// NewCell creates a new Cell instance.
// If id is empty, a new UUID will be generated.
func NewCell(workbook_id string, sheet_id string, row int, column string, value CellValue, id string) (*Cell, error) {
	if id == "" {
		var err error
		id, err = utils.GenerateUUID()
		if err != nil {
			return nil, fmt.Errorf("failed to generate UUID for Cell: %w", err)
		}
	}

	return &Cell{
		ID:         id,
		WorkbookID: workbook_id,
		SheetID:    sheet_id,
		Row:        row,
		Column:     column,
		Value:      value,
	}, nil
}

// GetPosition returns the position of the cell as a string,
// concatenating the column and row information.
func (c *Cell) GetPosition() string {
	return c.Column + fmt.Sprint(c.Row)
}
