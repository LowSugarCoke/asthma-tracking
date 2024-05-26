package repository

import (
	"log"
	"database/sql"
	"main/pkg/types"
)

// PostgresDailyInformationRepository represents a PostgreSQL implementation of UserRepository.
type PostgresDailyInformationRepository struct {
	DB *sql.DB
}

// NewPostgresDailyInformationRepository creates a new instance of PostgresDailyInformationRepository.
func NewPostgresDailyInformationRepository(db *sql.DB) *PostgresDailyInformationRepository {
	return &PostgresDailyInformationRepository{
		DB: db,
	}
}


// Create new daily information
func (r *PostgresDailyInformationRepository) CreateDailyInformation (d *types.DailyInformation) error {
	//  e.g. insert daily information into table
	log.Println("PostgresDailyInformationRepository Createusers")
	
	// Prepare the SQL statement
	stmt, err := r.DB.Prepare("INSERT INTO DailyInformation (user_id, timestamp, question_id, answer_id) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

    // Iterate over each InformationItem and insert into the database
    for _, info := range d.Information {

        // Execute the SQL statement with the data from the InformationItem struct
        _, err := stmt.Exec(d.UserID, d.Timestamp, info.Question, info.Answer)
        if err != nil {
            return err
        }

        log.Printf("Inserted daily information for user %s\n", d.UserID)
    }

    return nil
}