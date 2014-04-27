package game_engine

import (
    "requests"
)

type location struct {
    x int
    y int
}

func newLocation(x int, y int) location {
    return location{x: x, y: y}
}

func (location location) serialize() *requests.LocationStruct {
    return &requests.LocationStruct{
        X: location.x,
        Y: location.y}
}
