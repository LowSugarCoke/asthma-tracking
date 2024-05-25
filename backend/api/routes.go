package api

import (
    "net/http"
    "main/internal/dependencies"
)

func RegisterRoutes(dependencies *dependencies.Dependencies) http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /v1/users", dependencies.UserHandler.GetUsers)
    mux.HandleFunc("POST /v1/users", dependencies.UserHandler.CreateUser)

    mux.HandleFunc("POST /v1/daily_information", dependencies.DailyInformationHandler.CreateDailyInformation)

    return mux
}
