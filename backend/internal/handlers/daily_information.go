package handlers

import (
    "fmt"
    "encoding/json"
    "net/http"
    "main/pkg/types"
    "github.com/google/uuid"
    "main/internal/repository"
)


// DailyInformationHandler contains handler functions related to daily information.
type DailyInformationHandler struct {
    DailyInformationRepository repository.DailyInformationRepository
}

// NewDailyInformationHandler creates a new DailyInformationHandler with the given DailyInformationRepository.
func NewDailyInformationHandler(dailyInfoRepo repository.DailyInformationRepository) *DailyInformationHandler {
    return &DailyInformationHandler{DailyInformationRepository: dailyInfoRepo}
}


// CreateDailyInformation is the handler for creating a new daily information.
func (h *DailyInformationHandler) CreateDailyInformation(w http.ResponseWriter, r *http.Request) {
    
    var dailyInformation types.DailyInformation

    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()

    err := decoder.Decode(&dailyInformation)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validate the fields
    if dailyInformation.UserID == uuid.Nil || dailyInformation.Timestamp.IsZero() || len(dailyInformation.Information) == 0 {
    http.Error(w, "Invalid daily information payload", http.StatusBadRequest)
    return
    }

    // Call the DailyInformationRepository to create the daily information in the database
    err = h.DailyInformationRepository.CreateDailyInformation(&dailyInformation)
    if err != nil {
        // If there's an error creating the user in the database, return an internal server error response
        http.Error(w, "Failed to create daily information", http.StatusInternalServerError)
        return
    }

        
    fmt.Println("Created daily information")

    // Return an empty JSON response with status 201 (Created)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("{}"))
}
