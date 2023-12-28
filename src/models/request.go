package models

type AddPlayerRequest struct {
	Hash      string  `json:"hash"`
	Uuid      string  `json:"uuid"`
	Name      string  `json:"name"`
	Numbers   []int   `json:"numbers"`
	BingoCard [][]int `json:"bingoCard"`
}

type AddPlayerUsedCodeRequest struct {
	Hash string `json:"hash"`
	Code string `json:"code"`
	Uuid string `json:"uuid"`
}

type GameHasFinishedRequest struct {
	Hash string `json:"hash"`
	Uuid string `json:"uuid"`
}

type DeleteGameRequest struct {
	Hash string `json:"hash"`
	Uuid string `json:"uuid"`
}

type UpdateGameWinnersRequest struct {
	Hash    string            `json:"hash"`
	Uuid    string            `json:"uuid"`
	Winners map[string]string `json:"winners"`
}

type AddDrawnNumbersRequest struct {
	Uuid        string `json:"uuid"`
	DrawnNumber int    `json:"drawnNumber"`
}

type GetGameRequest struct {
	Hash string `json:"hash"`
}

type GetPlayerRequest struct {
	Hash string `json:"hash"`
	Uuid string `json:"uuid"`
}

type GetHostGameRequest struct {
	Uuid string `json:"uuid"`
}

type CreateGameRequest struct {
	Codes      *[]string `json:"codes"`
	Hash       string    `json:"hash"`
	Host       string    `json:"host"`
	MaxPlayers *int      `json:"maxPlayers"`
	Mode       string    `json:"mode"`
}
