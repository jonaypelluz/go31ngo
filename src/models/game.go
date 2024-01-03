package models

import "time"

type Game struct {
	Codes        []string          `bson:"codes" json:"codes"`
	DrawnNumbers []int             `bson:"drawnnumbers,omitempty" json:"drawnNumbers,omitempty"`
	HasFinished  bool              `bson:"hasfinished" json:"hasFinished"`
	HasStarted   bool              `bson:"hasstarted" json:"hasStarted"`
	Hash         string            `bson:"hash" json:"hash"`
	Host         string            `bson:"host" json:"host"`
	MaxPlayers   int               `bson:"maxplayers" json:"maxPlayers"`
	Mode         string            `bson:"mode" json:"mode"`
	Winners      map[string]string `bson:"winners,omitempty" json:"winners,omitempty"`
	Players      []string          `bson:"players,omitempty" json:"players,omitempty"`
	UsedCodes    []string          `bson:"usedcodes,omitempty" json:"usedCodes,omitempty"`
	CreatedAt    time.Time         `bson:"createdat" json:"createdAt"`
}
