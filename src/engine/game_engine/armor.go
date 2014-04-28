package game_engine

import (
    "requests"
)

type armorType int

const (
    cloth armorType = iota
    leather
    chainMail
    plate
)

type armor struct {
    name string
    armorType armorType
    strength int
}

func newArmor(name string, armorType armorType, strength int) *armor {
    return &armor{name: name, armorType: armorType, strength: strength}
}

func (armor *armor) serialize() *requests.ArmorStruct {
    return &requests.ArmorStruct{
        Name: armor.name,
        ArmorType: int(armor.armorType),
        Strength: armor.strength}
}
