package game

import (
	"api"
)

type armorType int

type armor struct {
	name      string
	armorType armorType
	strength  int
}

func newArmor(name string, armorType armorType, strength int) *armor {
	return &armor{name: name, armorType: armorType, strength: strength}
}

func (armor *armor) serialize() *api.ArmorStruct {
	return &api.ArmorStruct{
		Name:      armor.name,
		ArmorType: int(armor.armorType),
		Strength:  armor.strength}
}
