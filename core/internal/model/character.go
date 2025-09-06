package model

type Character struct {
	Id        string    `json:"id" redis:"id"`
	Equipment Equipment `json:"equipment" redis:"equipment"`
	Inventory Inventory `json:"inventory" redis:"inventory"`
}

type EquipmentType string

const (
	Head      EquipmentType = "head"
	Body      EquipmentType = "body"
	Legs      EquipmentType = "legs"
	LeftHand  EquipmentType = "left_hand"
	RightHand EquipmentType = "right_hand"
)

type Equipment struct {
	Items map[EquipmentType]Item `json:"items" redis:"items"`
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

func (stat *Stats) AddStat(newStat Stats) {
	stat.HP += newStat.HP
	stat.Aggro += newStat.Aggro
}
