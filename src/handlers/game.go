package handlers

import (
	"context"
	"go31ngo/src/models"
	"go31ngo/src/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateGameWinners(w http.ResponseWriter, r *http.Request, requestData models.UpdateGameWinnersRequest, dbService utils.DBService) {
	filter := bson.M{"host": requestData.Uuid, "hash": requestData.Hash}

	update := bson.M{
		"$set": bson.M{"winners": requestData.Winners},
	}

	result, err := dbService.UpdateOne(context.Background(), "games", filter, update)
	if err != nil {
		utils.SendApiResponse(w, "Error updating game: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		utils.SendApiResponse(w, "No matching game found or game already finished", http.StatusNoContent)
		return
	}

	utils.SendApiResponse(w, "Game winners updated successfully", http.StatusNoContent)
}

func AddDrawnNumber(w http.ResponseWriter, r *http.Request, requestData models.AddDrawnNumberRequest, dbService utils.DBService) {
	filter := bson.M{"host": requestData.Uuid, "hasfinished": false}

	update := bson.M{
		"$push": bson.M{"drawnnumbers": requestData.DrawnNumber},
		"$set":  bson.M{"hasstarted": true},
	}

	result, err := dbService.UpdateOne(context.Background(), "games", filter, update)
	if err != nil {
		utils.SendApiResponse(w, "Error updating game: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		utils.SendApiResponse(w, "No matching game found or game already finished", http.StatusNoContent)
		return
	}

	utils.SendApiResponse(w, "Drawn number added successfully", http.StatusNoContent)
}

func DeleteGame(w http.ResponseWriter, r *http.Request, requestData models.DeleteGameRequest, dbService utils.DBService) {
	gameFilter := bson.M{"host": requestData.Uuid, "hash": requestData.Hash}
	_, err := dbService.DeleteOne(context.Background(), "games", gameFilter)
	if err != nil {
		utils.SendApiResponse(w, "Error deleting game: "+err.Error(), http.StatusInternalServerError)
		return
	}

	playerFilter := bson.M{"hash": requestData.Hash}
	_, err = dbService.DeleteMany(context.Background(), "players", playerFilter)
	if err != nil {
		utils.SendApiResponse(w, "Error deleting associated players: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendApiResponse(w, "Game deleted successfully", http.StatusNoContent)
}

func GameHasFinished(w http.ResponseWriter, r *http.Request, requestData models.GameHasFinishedRequest, dbService utils.DBService) {
	filter := bson.M{"host": requestData.Uuid, "hash": requestData.Hash}

	update := bson.M{
		"$set": bson.M{"hasfinished": true},
	}

	result, err := dbService.UpdateOne(context.Background(), "games", filter, update)
	if err != nil {
		utils.SendApiResponse(w, "Error updating game: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		utils.SendApiResponse(w, "No matching game found or game already finished", http.StatusNoContent)
		return
	}

	utils.SendApiResponse(w, "Game finish status changed successfully", http.StatusNoContent)
}

func GetGame(w http.ResponseWriter, r *http.Request, requestData models.GetGameRequest, dbService utils.DBService) {
	filter := bson.M{"hash": requestData.Hash}

	result, err := dbService.FindBy(context.Background(), "games", filter)
	if err != nil {
		utils.SendApiResponse(w, "Error retrieving game: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var game models.Game
	if err := result.Decode(&game); err != nil {
		if err == mongo.ErrNoDocuments {
			utils.SendApiResponse(w, "No game found", http.StatusNoContent)
			return
		} else {
			utils.SendApiResponse(w, "Error decoding game data: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	utils.SendApiResponse(w, game, http.StatusOK)
}

func GetHostGame(w http.ResponseWriter, r *http.Request, requestData models.GetHostGameRequest, dbService utils.DBService) {
	filter := bson.M{"host": requestData.Uuid}

	result, err := dbService.FindBy(context.Background(), "games", filter)
	if err != nil {
		utils.SendApiResponse(w, "Error retrieving game: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var game models.Game
	if err := result.Decode(&game); err != nil {
		if err == mongo.ErrNoDocuments {
			utils.SendApiResponse(w, "No game found", http.StatusNoContent)
			return
		} else {
			utils.SendApiResponse(w, "Error decoding game data: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	utils.SendApiResponse(w, game, http.StatusOK)
}

func AddPlayerUsedCode(w http.ResponseWriter, r *http.Request, requestData models.AddPlayerUsedCodeRequest, dbService utils.DBService) {
	filter := bson.M{
		"hash":      requestData.Hash,
		"codes":     requestData.Code,
		"usedcodes": bson.M{"$ne": requestData.Code},
	}

	update := bson.M{
		"$push": bson.M{"usedcodes": requestData.Code, "players": requestData.Uuid},
	}

	result, err := dbService.UpdateOne(context.Background(), "games", filter, update)
	if err != nil {
		utils.SendApiResponse(w, "Error updating game: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		utils.SendApiResponse(w, "Code not valid or already used", http.StatusConflict)
		return
	}

	utils.SendApiResponse(w, "Code added successfully", http.StatusNoContent)
}

func CreateGame(w http.ResponseWriter, r *http.Request, requestData models.CreateGameRequest, dbService utils.DBService) {
	codes := []string{}
	if requestData.Codes != nil {
		codes = *requestData.Codes
	}

	maxPlayers := 0
	if requestData.MaxPlayers != nil {
		maxPlayers = *requestData.MaxPlayers
	}

	newGame := models.Game{
		Codes:        codes,
		DrawnNumbers: []int{},
		HasFinished:  false,
		HasStarted:   false,
		Hash:         requestData.Hash,
		Host:         requestData.Host,
		MaxPlayers:   maxPlayers,
		Mode:         requestData.Mode,
		Winners:      make(map[string]string),
		Players:      []string{},
		UsedCodes:    []string{},
		CreatedAt:    time.Now(),
	}

	_, err := dbService.InsertOne(context.Background(), "games", newGame)
	if err != nil {
		utils.SendApiResponse(w, "Error creating game", http.StatusInternalServerError)
		return
	}

	utils.SendApiResponse(w, "Game created successfully", http.StatusOK)
}
