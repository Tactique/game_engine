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
    armor *armor
    canAttack bool
}

func newUnit(name string, nation nation, movement *movement, attacks []*attack, armor *armor) *unit {
    return &unit{
        name: name, health: 10, nation: nation,
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
    sword := newAttack("Basic Sword", sword, 2)
    chainMailArmor := newArmor("Chain Mail", chainMail, 3)
    return newUnit(
        "Warrior",
        nation,
        legs,
        []*attack{sword},
        chainMailArmor)
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
        CanMove: unit.canMove,
        CanAttack: unit.canAttack,
        Attacks: attacks,
        Armor: unit.armor.serialize()}
}

func (unit *unit) turnReset() {
    unit.canMove = true
    unit.canAttack = true
}
