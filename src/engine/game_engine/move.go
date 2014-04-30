package game_engine

import (
	"errors"
	"fmt"
)

func getDistance(first int, second int) int {
	if (first - second) > 0 {
		return first - second
	} else {
		return second - first
	}
}

func assertAllTilesTouch(locations []location) error {
	oldLoc := locations[0]
	for _, loc := range locations[1:] {
		if (getDistance(loc.x, oldLoc.x) + getDistance(loc.y, oldLoc.y)) != 1 {
			fmt.Println("Two units too far away: ")
			return errors.New(fmt.Sprintf(
				"Two units too far away (%d, %d) (%d, %d)",
				loc.x, oldLoc.x, loc.y, oldLoc.y))
		}
		oldLoc = loc
	}
	return nil
}

func validMove(moves int, movementType *movement, tiles []terrain, locations []location) error {
	if len(locations) < 2 {
		return errors.New("Need at least two tiles, a start and an end")
	}
	err := assertAllTilesTouch(locations)
	if err != nil {
		return err
	}
	for _, tile := range tiles {
		moves -= int(movementType.costs[tile])
	}
	fmt.Println("Had this many moves left", moves)
	if moves >= 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Too much distance to cover, need %d less", moves))
	}
}
