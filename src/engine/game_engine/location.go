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

func locationFromRequest(requestLocation requests.LocationStruct) location {
    return newLocation(requestLocation.X, requestLocation.Y)
}

func (location location) serialize() *requests.LocationStruct {
    return &requests.LocationStruct{
        X: location.x,
        Y: location.y}
}
