package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/Deathfireofdoom/terraxcel/server/pkg/db"
)

type DocumentRepository struct {
	dbManager *db.DBManager
}

func NewDocumentRepository() (*DocumentRepository, error) {
	log.Println("creating repository")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" {
		dbHost = "localhost"
	}

	if dbPort == "" {
		dbPort = "5432"
	}

	if dbUser == "" {
		dbUser = "postgres"
	}

	if dbPassword == "" {
		dbPassword = "postgres"
	}

	if dbName == "" {
		dbName = "postgres"
	}

	dataSourceName := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	dbManager, err := db.NewDBManager(dataSourceName)
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
