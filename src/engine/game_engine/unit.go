package game_engine

import (
    "requests"
)

type unit struct {
    name string
    health int
    nation nation
    movement *movement
    canMove bool
    attacks []*attack
    canAttack bool
}

func newUnit(name string, nation nation, movement *movement, attacks []*attack) *unit {
    return &unit{
        name: name, health: 10, nation: nation,
        movement: movement, canMove: true,
        attacks: attacks, canAttack: true}
}

func warrior(nation nation) *unit {
    legs := newMovement(
        "Legs",
        10,
        map[terrain]multiplier{
            plains: multiplier(1.0)})
    sword := newAttack("Basic Sword", sword, 5)
    return newUnit(
        "Warrior",
        nation,
        legs,
        []*attack{sword})
}

func (unit *unit) serialize(loc location) *requests.UnitStruct {
    attacks := make([]*requests.AttackStruct, len(unit.attacks))
    for i, attack := range(unit.attacks) {
        attacks[i] = attack.serialize()
    }
    return &requests.UnitStruct{
        Name: unit.name,
        Health: unit.health,
        Nation: int(unit.nation),
        Movement: unit.movement.serialize(),
        Position: loc.serialize(),
        Distance: unit.movement.distance,
        CanMove: unit.canMove,
        CanAttack: unit.canAttack,
        Attacks: attacks}
}

func (unit *unit) turnReset() {
    unit.canMove = true
    unit.canAttack = true
}
