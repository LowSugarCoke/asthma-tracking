package handlers

import (
    "fmt"
    "encoding/json"
    "net/http"
    "main/pkg/types"
    "github.com/google/uuid"
    "main/internal/repository"
    
)


// UsersHandler contains handler functions related to users.
type UsersHandler struct {
    UserRepository repository.UserRepository
}

// NewUsersHandler creates a new UsersHandler with the given UserRepository.
func NewUsersHandler(userRepo repository.UserRepository) *UsersHandler {
    return &UsersHandler{UserRepository: userRepo}
}


// CreateUserHandler handles requests to create a new user.
func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user types.User

    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()

    err := decoder.Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	// Todo: Validate user fields
	if !ValidateUser(user) {
		return
	}

    // Generate a UUID for the ID field
	user.ID = generateUUID()

    fmt.Printf("Decoded User: %+v\n", user)

    // Call the UserRepository to create the user in the database
    err = h.UserRepository.CreateUser(&user)
    if err != nil {
        // If there's an error creating the user in the database, return an internal server error response
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    // Return an empty JSON response with status 201 (Created)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("{}"))
    
}

// GetUserHandler handles requests to retrieve a user by ID.
func (h *UsersHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

        
    fmt.Println("Getting users")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Getting users"))
}




// generateUUID generates a UUID for the user ID.
func generateUUID() uuid.UUID {
	// "123e4567-e89b-12d3-a456-426614174000" // Dummy UUID for demonstration
    
    // Generate a new UUID
    newUUID := uuid.New()
    return newUUID
}