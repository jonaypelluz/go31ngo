package main

import (
	"fmt"
	"go31ngo/src/config"
	"go31ngo/src/handlers"
	"go31ngo/src/middleware"
	"go31ngo/src/utils"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, cfg *config.Config, mongoService *utils.MongoService) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, ":-)")
	})

	apiPrefix := func(pattern string, handler http.Handler) {
		prefixedPattern := cfg.APIVersion + pattern
		mux.Handle(prefixedPattern, handler)
	}

	apiPrefix("/add-player", middleware.RequestCheck("POST", handlers.AddPlayer, mongoService))
	apiPrefix("/get-player", middleware.RequestCheck("POST", handlers.GetPlayer, mongoService))
	apiPrefix("/add-player-used-code", middleware.RequestCheck("POST", handlers.AddPlayerUsedCode, mongoService))
	apiPrefix("/add-drawn-numbers", middleware.RequestCheck("POST", handlers.AddDrawnNumbers, mongoService))
	apiPrefix("/create-game", middleware.RequestCheck("POST", handlers.CreateGame, mongoService))
	apiPrefix("/get-host-game", middleware.RequestCheck("POST", handlers.GetHostGame, mongoService))
	apiPrefix("/has-finished", middleware.RequestCheck("POST", handlers.GameHasFinished, mongoService))
	apiPrefix("/get-game", middleware.RequestCheck("POST", handlers.GetGame, mongoService))
	apiPrefix("/update-winners", middleware.RequestCheck("POST", handlers.UpdateGameWinners, mongoService))
	apiPrefix("/delete-game", middleware.RequestCheck("POST", handlers.DeleteGame, mongoService))
}
