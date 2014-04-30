package game_engine

import (
	"api"
	"errors"
)

type unit struct {
	name      string
	health    int
	maxHealth int
	nation    nation
	movement  *movement
	canMove   bool
	attacks   []*attack
	armor     *armor
	canAttack bool
}

func newUnit(name string, nation nation, movement *movement, attacks []*attack, armor *armor) *unit {
	return &unit{
		name: name, health: 10, maxHealth: 10,
		nation:   nation,
		movement: movement, canMove: true,
		attacks: attacks, canAttack: true,
		armor: armor}
}

func warrior(nation nation) *unit {
	legs := newMovement(
		"Legs",
		10,
		map[terrain]multiplier{
			plains: multiplier(1.0)})
	sword := newAttack("Basic Sword", sword, 4)
	chainMailArmor := newArmor("Chain Mail", chainMail, 2)
	return newUnit(
		"Warrior",
		nation,
		legs,
		[]*attack{sword},
		chainMailArmor)
}

func (unit *unit) serialize(loc location) *api.UnitStruct {
	attacks := make([]*api.AttackStruct, len(unit.attacks))
	for i, attack := range unit.attacks {
		attacks[i] = attack.serialize()
	}
	return &api.UnitStruct{
		Name:      unit.name,
		Health:    unit.health,
		Nation:    int(unit.nation),
		Movement:  unit.movement.serialize(),
		Position:  loc.serialize(),
		CanMove:   unit.canMove,
		CanAttack: unit.canAttack,
		Attacks:   attacks,
		Armor:     unit.armor.serialize()}
}

func (unit *unit) turnReset() {
	unit.canMove = true
	unit.canAttack = true
}

func (unit *unit) receiveDamage(delta int) (bool, error) {
	if delta < 0 {
		return true, errors.New("Cannot receive a negative amount of damage")
	}
	return unit.changeHealth(unit.health - delta), nil
}

func (unit *unit) changeHealth(newHealth int) bool {
	unit.health = newHealth
	if unit.health > unit.maxHealth {
		unit.health = unit.maxHealth
	}
	if unit.health <= 0 {
		unit.health = 0
		return false
	}
	return true
}
