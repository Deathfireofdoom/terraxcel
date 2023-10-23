package repository

import "log"

type DocumentRepository struct {
	dbManager *DBManager
}

func NewDocumentRepository() (*DocumentRepository, error) {
	dbManager, err := NewDBManager("user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Printf("failed to create db manager: %v", err)
		return nil, err
	}

	repository := &DocumentRepository{dbManager}
	err = repository.initialize()
	if err != nil {
		log.Printf("failed to initialize repository: %v", err)
		return nil, err
	}

	return repository, nil
}

// initialize creates the necessary tables in the database.
func (r *DocumentRepository) initialize() error {
	log.Println("initializing repository")

	log.Println("creating workbook table")
	if err := r.createWorkbookTable(); err != nil {
		log.Printf("failed to create workbook table: %v", err)
		return err
	}

	log.Println("creating sheet table")
	if err := r.createSheetTable(); err != nil {
		log.Printf("failed to create sheet table: %v", err)
		return err
	}

	log.Println("creating cell table")
	if err := r.createCellTable(); err != nil {
		log.Printf("failed to create cell table: %v", err)
		return err
	}

	log.Println("repository initialized")
	return nil
}
