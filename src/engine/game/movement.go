package game

import (
	"api"
	"strconv"
)

type multiplier float64

type Movement struct {
	name     string
	distance int
	costs    map[Terrain]multiplier
}

func NewMovement(name string, distance int, costs map[Terrain]multiplier) *Movement {
	return &Movement{name: name, distance: distance, costs: costs}
}

func (movement *Movement) GetDistance() int {
	return movement.distance
}

func (movement *Movement) serialize() *api.MovementStruct {
	costs := make(map[string]float64, 0)
	for terrain, cost := range movement.costs {
		costs[strconv.Itoa(int(terrain))] = float64(cost)
	}
	return &api.MovementStruct{
		Type:     movement.name,
		Speeds:   costs,
		Distance: movement.distance}
}
