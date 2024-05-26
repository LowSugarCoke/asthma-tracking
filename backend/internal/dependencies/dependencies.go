package dependencies

import "main/internal/handlers"

// Dependencies struct holds the dependencies needed by the server.
type Dependencies struct {
    UserHandler *handlers.UsersHandler
    DailyInformationHandler *handlers.DailyInformationHandler
}

// NewDependencies creates and initializes the dependencies.
func NewDependencies(userHandler *handlers.UsersHandler, dailyInfoHandler *handlers.DailyInformationHandler) *Dependencies {
    return &Dependencies{
        UserHandler: userHandler,
        DailyInformationHandler: dailyInfoHandler,
    }
}
