package game_engine

import (
	"errors"
	"fmt"
	"github.com/Tactique/golib/logger"
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
		distance := (getDistance(loc.x, oldLoc.x) + getDistance(loc.y, oldLoc.y))
		if distance != 1 {
			locations := fmt.Sprintf(
				"(%d, %d) (%d, %d)", loc.x, oldLoc.x, loc.y, oldLoc.y)
			logger.Infof("Two units too far away, distance %d (locations %s)", distance, locations)
			return errors.New(fmt.Sprintf("Two units too far away %s", locations))
		}
		oldLoc = loc
	}
	return nil
}

func validMove(moves int, movementType *movement, tiles []terrain, locations []location) error {
	if len(locations) < 2 {
		message := "Need at least two tiles, a start and an end"
		logger.Infof(message + " (got %d)", len(locations))
		return errors.New(message)
	}
	err := assertAllTilesTouch(locations)
	if err != nil {
		return err
	}
	for _, tile := range tiles {
		moves -= int(movementType.costs[tile])
	}
	logger.Debugf("Had this many moves left %d", moves)
	if moves >= 0 {
		return nil
	} else {
		message := fmt.Sprintf("Too much distance to cover, need %d less", moves)
		logger.Infof(message)
		return errors.New(message)
	}
}
