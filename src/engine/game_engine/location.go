package game_engine

import (
    "api"
)

type location struct {
    x int
    y int
}

func newLocation(x int, y int) location {
    return location{x: x, y: y}
}

func locationFromRequest(requestLocation api.LocationStruct) location {
    return newLocation(requestLocation.X, requestLocation.Y)
}

func (location location) serialize() *api.LocationStruct {
    return &api.LocationStruct{
        X: location.x,
        Y: location.y}
}
