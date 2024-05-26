package types

import (
    "github.com/google/uuid"
    "time"
)


// Original
// InformationItem represents a question and its answer.
// type InformationItem struct {
//     Question string `json:"question"`
//     Answer   string `json:"answer"`
// }


// Suggested
type InformationItem struct {
    Question int `json:"question"`
    Answer   int `json:"answer"`
}

// DailyInformation represents the daily information for a user.
type DailyInformation struct {
    UserID     uuid.UUID `json:"user_id"`
    Timestamp  time.Time `json:"timestamp"`
    Information []InformationItem `json: "information"`
}

type User struct {
	ID       uuid.UUID    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}