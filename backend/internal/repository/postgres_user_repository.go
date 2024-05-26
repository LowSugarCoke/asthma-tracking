package repository

import (
	"log"
	"database/sql"
	"main/pkg/types"
)

// PostgresUserRepository represents a PostgreSQL implementation of UserRepository.
type PostgresUserRepository struct {
	DB *sql.DB
}

// NewPostgresUserRepository creates a new instance of PostgresUserRepository.
func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		DB: db,
	}
}

// GetAllUsers retrieves all users from the database.
func (r *PostgresUserRepository) GetAllUsers() ([]types.User, error) {
	// Implement logic to fetch users from the database using r.DB
	log.Println("PostgresUserRepository GetAllusers")
	
	return nil, nil
}

// Create new user 
func (r *PostgresUserRepository) CreateUser (u *types.User) error {
	//  e.g. insert user into table
	log.Println("PostgresUserRepository Createusers")
	
	// Prepare the SQL statement
	stmt, err := r.DB.Prepare("INSERT INTO Users (id, name, email) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the data from the InformationItem struct
	_, err = stmt.Exec(u.ID, u.Username, u.Email)
	if err != nil {
		return err
	}

	log.Printf("Created new user %s\n", u.ID)

	return nil
}