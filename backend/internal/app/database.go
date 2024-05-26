package app

import (
	"os"
	"fmt"
    "log"
	"time"
    "io/ioutil"
    "database/sql"
    _ "github.com/lib/pq"
)


func ConnectDB() (*sql.DB, error) {

	var db *sql.DB
	var err error

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

	// Todo: In production mode, change sslmode to require or verify-full
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
    
	fmt.Println("Database connection info:", psqlInfo) 

	// Retry logic
    for i := 0; i < 10; i++ {
        db, err = sql.Open("postgres", psqlInfo)
        if err == nil {
            err = db.Ping()
            if err == nil {
                break
            }
        }
        fmt.Println("Retrying to connect to the database...")
        time.Sleep(2 * time.Second)
    }

    if err != nil {
		return nil, err
    }


    // Initialize schema
    if err := initializeSchema(db); err != nil {
        return nil, err
    }

    return db, nil
}


func initializeSchema(db *sql.DB) error {
    schemaFile := "sql/schema.sql"
    schema, err := ioutil.ReadFile(schemaFile)
    if err != nil {
        return fmt.Errorf("Could not read schema file: %v", err)
    }

    _, err = db.Exec(string(schema))
    if err != nil {
        return fmt.Errorf("could not execute schema: %v", err)
    }

    log.Println("Database schema initialized successfully")
    return nil
}