package repository

import "main/pkg/types"

// DailyInformationRepository defines the methods a user repository must implement.
type DailyInformationRepository interface {
	CreateDailyInformation(dailyInformation *types.DailyInformation) error
}

