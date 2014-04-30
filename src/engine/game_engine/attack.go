package game_engine

import (
    "api"
)

type attackType int

const (
    sword attackType = iota
    axe
    mace
    wand
    staff
)

type attack struct {
    name string
    attackType attackType
    power int
}

func newAttack(name string, attackType attackType, power int) *attack {
    return &attack{name: name, attackType: attackType, power: power}
}

func (attack *attack) serialize() *api.AttackStruct {
    return &api.AttackStruct{
        Name: attack.name,
        AttackType: int(attack.attackType),
        Power: attack.power}
}
