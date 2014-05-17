package game

import (
	"api"
)

type attackType int

type attack struct {
	name       string
	attackType attackType
	power      int
	minRange   int
	maxRange   int
}

func newAttack(name string, attackType attackType, power int, minRange int, maxRange int) *attack {
	return &attack{
		name:       name,
		attackType: attackType,
		power:      power,
		minRange:   minRange,
		maxRange:   maxRange,
	}
}

func (attack *attack) serialize() *api.AttackStruct {
	return &api.AttackStruct{
		Name:       attack.name,
		AttackType: int(attack.attackType),
		Power:      attack.power,
		MinRange:   attack.minRange,
		MaxRange:   attack.maxRange,
	}
}
