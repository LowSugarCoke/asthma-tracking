package main

import (
    "log" 
    "main/internal/app"
)



func main() {    
    app.Load()

    db, err := app.ConnectDB()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    defer db.Close()
    log.Println("Successfully connected!")

    port := ":8080"


    // Create and start the server
    server := app.NewServer(db, port)
    server.Start()

}

