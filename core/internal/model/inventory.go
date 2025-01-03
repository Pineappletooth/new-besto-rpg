package model

type Inventory struct {
	Items map[string]int `json:"items" redis:"items"`
	Gold  int64            `json:"gold" redis:"gold"`
}