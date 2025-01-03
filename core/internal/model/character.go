package model

type Character struct {
	Id        uint32    `json:"id" redis:"id"`
	Inventory Inventory `json:"inventory" redis:"inventory"`
}