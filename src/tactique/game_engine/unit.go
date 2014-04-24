package game_engine

import (
    "requests"
)

type unit struct {
    name string
    health int
    nation nation
    movement *movement
}

func newUnit(name string, nation nation, movement *movement) *unit {
    return &unit{name: name, health: 10, nation: nation, movement: movement}
}

func tank(nation nation) *unit {
    return newUnit(
        "Tank",
        nation,
        newMovement(
            "Treads",
            10,
            map[terrain]multiplier{
                plains: multiplier(1.0)}))
}

func (unit *unit) serialize(loc location) *requests.UnitStruct {
    return &requests.UnitStruct{
        Name: unit.name,
        Health: unit.health,
        Nation: int(unit.nation),
        Movement: unit.movement.serialize(),
        Position: loc.serialize(),
        Distance: unit.movement.distance}
}
