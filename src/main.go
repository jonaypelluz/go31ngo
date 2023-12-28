package main

import (
	"fmt"
	"go31ngo/src/handlers"
	"go31ngo/src/middleware"
	"go31ngo/src/utils"
	"net/http"
	"os"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")

	utils.ConnectMongoDB(mongoURI)
	mongoService := utils.TheMongoService(utils.GetMongoDBClient())

	apiVersion := "/api/v1"

	http.Handle(apiVersion+"/add-player", middleware.RequestCheck("POST", handlers.AddPlayer, mongoService))
	http.Handle(apiVersion+"/get-player", middleware.RequestCheck("POST", handlers.GetPlayer, mongoService))
	http.Handle(apiVersion+"/add-player-used-code", middleware.RequestCheck("POST", handlers.AddPlayerUsedCode, mongoService))
	http.Handle(apiVersion+"/add-drawn-numbers", middleware.RequestCheck("POST", handlers.AddDrawnNumbers, mongoService))
	http.Handle(apiVersion+"/create-game", middleware.RequestCheck("POST", handlers.CreateGame, mongoService))
	http.Handle(apiVersion+"/get-host-game", middleware.RequestCheck("POST", handlers.GetHostGame, mongoService))
	http.Handle(apiVersion+"/has-finished", middleware.RequestCheck("POST", handlers.GameHasFinished, mongoService))
	http.Handle(apiVersion+"/get-game", middleware.RequestCheck("POST", handlers.GetGame, mongoService))
	http.Handle(apiVersion+"/update-winners", middleware.RequestCheck("POST", handlers.UpdateGameWinners, mongoService))
	http.Handle(apiVersion+"/delete-game", middleware.RequestCheck("POST", handlers.DeleteGame, mongoService))

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
