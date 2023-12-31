package handlers

import (
	"context"
	"go31ngo/src/models"
	"go31ngo/src/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddPlayer(w http.ResponseWriter, r *http.Request, requestData models.AddPlayerRequest, dbService utils.DBService) {
	filter := bson.M{"uuid": requestData.Uuid, "hash": requestData.Hash}

	existingPlayer, err := dbService.FindBy(context.Background(), "players", filter)
	if err != nil && err != mongo.ErrNoDocuments {
		utils.SendApiResponse(w, "Error checking existing player info", http.StatusInternalServerError)
		return
	}

	if existingPlayer != nil && existingPlayer.Err() != mongo.ErrNoDocuments {
		utils.SendApiResponse(w, "Player already exists", http.StatusConflict)
		return
	}

	_, err = dbService.InsertOne(context.Background(), "players", requestData)
	if err != nil {
		utils.SendApiResponse(w, "Error adding player info", http.StatusInternalServerError)
		return
	}

	utils.SendApiResponse(w, "", http.StatusNoContent)
}

func GetPlayer(w http.ResponseWriter, r *http.Request, requestData models.GetPlayerRequest, dbService utils.DBService) {
	filter := bson.M{"uuid": requestData.Uuid, "hash": requestData.Hash}

	result, err := dbService.FindBy(context.Background(), "players", filter)
	if err != nil {
		utils.SendApiResponse(w, "Error retrieving players info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var player models.Player
	if err := result.Decode(&player); err != nil {
		if err == mongo.ErrNoDocuments {
			utils.SendApiResponse(w, "No player info found", http.StatusNoContent)
			return
		} else {
			utils.SendApiResponse(w, "Error decoding player info data: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	utils.SendApiResponse(w, player, http.StatusOK)
}
