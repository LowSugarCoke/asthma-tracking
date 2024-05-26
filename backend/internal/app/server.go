package app

import (
    "log"
    "time"
    "net/http"
    "database/sql"
    "main/api"
	"main/internal/handlers"
	"main/internal/repository"
	"main/internal/dependencies"
)


type Server struct {
    // DB   *sql.DB     // keep this field if I need direct access to the database connection at the server level (outside of the repository layer)
    Port string
	Dependencies *dependencies.Dependencies
}


func NewServer(db *sql.DB, port string) *Server {
	
	// Initialize repositories and create dependencies
	userRepo := &repository.PostgresUserRepository{DB: db}
    dailyInfoRepo := &repository.PostgresDailyInformationRepository{DB: db}
	dependencies := dependencies.NewDependencies(handlers.NewUsersHandler(userRepo), handlers.NewDailyInformationHandler(dailyInfoRepo))

	return &Server{
		Port:         port,
		Dependencies: dependencies,
	}
}



func (s *Server) Start() {

    // Initialize the router
    router := api.RegisterRoutes(s.Dependencies)

    // Custom HTTP server configuration
    server := &http.Server{
        Addr:           s.Port,
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    // Start the server
    log.Printf("Server is running on http://localhost%s\n", s.Port)
    if err := server.ListenAndServe(); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}
