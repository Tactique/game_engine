package game_engine

import (
    "api"
    "strconv"
)

type multiplier float64

type movement struct {
    name string
    distance int
    costs map[terrain]multiplier
}

func newMovement(name string, distance int, costs map[terrain]multiplier) *movement {
    return &movement{name: name, distance: distance, costs: costs}
}

func (movement *movement) serialize() *api.MovementStruct {
    costs := make(map[string]float64, 0)
    for terrain, cost := range(movement.costs) {
        costs[strconv.Itoa(int(terrain))] = float64(cost)
    }
    return &api.MovementStruct{
        Type: movement.name,
        Speeds: costs,
        Distance: movement.distance}
}
