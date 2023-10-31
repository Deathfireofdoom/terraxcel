package repository

import "github.com/Deathfireofdoom/terraxcel/common/models"

func (r *DocumentRepository) createWorkbookTable() error {
	tableQuery := `
		CREATE TABLE IF NOT EXISTS workbook (
			id 			VARCHAR(255) PRIMARY KEY,
			file_name 	VARCHAR(255) NOT NULL,
			folder_path VARCHAR(255) NOT NULL,
			extension 	VARCHAR(255) NOT NULL
			last_modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := r.dbManager.Exec(tableQuery)

	if err != nil {
		return err
	}

	return nil
}

func (r *DocumentRepository) GetWorkbook(id string) (*models.Workbook, error) {
	query := `
		SELECT * FROM workbook WHERE id = $1;
	`

	row := r.dbManager.QueryRow(query, id)

	var workbook models.Workbook
	err := row.Scan(
		&workbook.ID,
		&workbook.FileName,
		&workbook.FolderPath,
		&workbook.Extension,
	)

	if err != nil {
		return nil, err
	}

	return &workbook, nil
}

func (r *DocumentRepository) SaveWorkbook(workbook *models.Workbook) error {
	query := `
		INSERT INTO workbook (
			id,
			file_name,
			folder_path,
			extension
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) ON CONFLICT (id) DO UPDATE SET
			file_name = $2,
			folder_path = $3,
			extension = $4
			last_modified = CURRENT_TIMESTAMP;
	`

	_, err := r.dbManager.Exec(
		query,
		workbook.ID,
		workbook.FileName,
		workbook.FolderPath,
		workbook.Extension,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *DocumentRepository) DeleteWorkbook(id string) error {
	query := `
		DELETE FROM workbook WHERE id = $1;
	`

	_, err := r.dbManager.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
