package repository

import "github.com/Deathfireofdoom/terraxcel/common/models"

func (r *DocumentRepository) createCellTable() error {
	tableQuery := `
		CREATE TABLE IF NOT EXISTS cell (
			id 			VARCHAR(255) PRIMARY KEY,
			sheet_id 	VARCHAR(255) NOT NULL,
			workbook_id VARCHAR(255) NOT NULL,
			row     	INT NOT NULL,
			col         VARCHAR(255) NOT NULL,
			value_type 	VARCHAR(255) NOT NULL,
			value 		VARCHAR(255) NOT NULL,
			last_modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := r.dbManager.Exec(tableQuery)

	if err != nil {
		return err
	}

	return nil
}

func (r *DocumentRepository) GetCell(id string) (*models.Cell, error) {
	query := `
		SELECT * FROM cell WHERE id = $1;
	`

	row := r.dbManager.QueryRow(query, id)

	var cell models.Cell
	err := row.Scan(
		&cell.ID,
		&cell.SheetID,
		&cell.WorkbookID,
		&cell.Row,
		&cell.Column,
		&cell.Value.Type,
		&cell.Value.Value,
	)

	if err != nil {
		return nil, err
	}

	return &cell, nil
}

func (r *DocumentRepository) SaveCell(cell *models.Cell) error {
	query := `
		INSERT INTO cell (
			id,
			sheet_id,
			workbook_id,
			row,
			col,
			value_type,
			value
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		) ON CONFLICT (id) DO UPDATE SET
			sheet_id = $2,
			workbook_id = $3,
			row = $4,
			col = $5,
			value_type = $6,
			value = $7
	`

	_, err := r.dbManager.Exec(
		query,
		cell.ID,
		cell.SheetID,
		cell.WorkbookID,
		cell.Row,
		cell.Column,
		cell.Value.Type,
		cell.Value.Value,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *DocumentRepository) DeleteCell(id string) error {
	query := `
		DELETE FROM cell WHERE id = $1;
	`

	_, err := r.dbManager.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
