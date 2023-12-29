package main

import (
	"fmt"
	"go31ngo/src/config"
	"go31ngo/src/utils"
	"net/http"
)

func main() {
	cfg := config.AppConfig()

	utils.ConnectMongoDB(cfg.MongoURI)
	mongoService := utils.TheMongoService(utils.GetMongoDBClient())

	mux := http.NewServeMux()

	SetupRoutes(mux, cfg, mongoService)

	fmt.Println("Server is running on port 8090...")
	http.ListenAndServe(":8090", mux)
}
