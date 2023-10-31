package repository

import "github.com/Deathfireofdoom/terraxcel/common/models"

func (r *DocumentRepository) createSheetTable() error {
	tableQuery := `
		CREATE TABLE IF NOT EXISTS sheet (
			id 			VARCHAR(255) PRIMARY KEY,
			workbook_id VARCHAR(255) NOT NULL,
			name 		VARCHAR(255) NOT NULL,
			pos 		INT NOT NULL,
			last_modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := r.dbManager.Exec(tableQuery)

	if err != nil {
		return err
	}

	return nil
}

func (r *DocumentRepository) GetSheet(id string) (*models.Sheet, error) {
	query := `
		SELECT * FROM sheet WHERE id = $1;
	`

	row := r.dbManager.QueryRow(query, id)

	var sheet models.Sheet
	err := row.Scan(
		&sheet.ID,
		&sheet.WorkbookID,
		&sheet.Name,
		&sheet.Pos,
	)

	if err != nil {
		return nil, err
	}

	return &sheet, nil
}

func (r *DocumentRepository) SaveSheet(sheet *models.Sheet) error {
	query := `
		INSERT INTO sheet (
			id,
			workbook_id,
			name,
			pos
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) ON CONFLICT (id) DO UPDATE SET
			workbook_id = $2,
			name = $3,
			pos = $4,
			last_modified = CURRENT_TIMESTAMP;
	`

	_, err := r.dbManager.Exec(
		query,
		sheet.ID,
		sheet.WorkbookID,
		sheet.Name,
		sheet.Pos,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *DocumentRepository) DeleteSheet(id string) error {
	query := `
		DELETE FROM sheet WHERE id = $1;
	`

	_, err := r.dbManager.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
