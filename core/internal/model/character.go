package model

type Character struct {
	Id        string    `json:"id" redis:"id"`
	Equipment Equipment `json:"equipment" redis:"equipment"`
	Inventory Inventory `json:"inventory" redis:"inventory"`
}

type Equipment struct {
	Head      Item `json:"head" redis:"head"`
	Body      Item `json:"body" redis:"body"`
	Legs      Item `json:"legs" redis:"legs"`
	LeftHand  Item `json:"left_hand" redis:"left_hand"`
	RightHand Item `json:"right_hand" redis:"right_hand"`
}

type Item struct {
	Name   string   `json:"name" redis:"name"`
	Skills []string `json:"skills" redis:"skills"`
	Stats  Stats    `json:"stats" redis:"stats"`
}

type Stat string

const (
	HP = "HP"
)

type Stats struct {
	HP    int
	Aggro int
}
