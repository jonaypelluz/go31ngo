package models

import "time"

type Player struct {
	Name      string    `bson:"name" json:"name"`
	Hash      string    `bson:"hash" json:"hash"`
	Numbers   []int     `bson:"numbers" json:"numbers"`
	BingoCard [][]int   `bson:"bingocard" json:"bingoCard"`
	Uuid      string    `bson:"uuid" json:"uuid"`
	CreatedAt time.Time `bson:"createdat" json:"createdAt"`
}
