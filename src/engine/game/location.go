package game

import (
	"api"
)

type Location struct {
	x int
	y int
}

func NewLocation(x int, y int) Location {
	return Location{x: x, y: y}
}

func LocationFromRequest(requestLocation *api.LocationStruct) Location {
	return NewLocation(requestLocation.X, requestLocation.Y)
}

func (location Location) GetX() int {
	return location.x
}

func (location Location) GetY() int {
	return location.y
}

func (location Location) serialize() *api.LocationStruct {
	return &api.LocationStruct{
		X: location.x,
		Y: location.y}
}
