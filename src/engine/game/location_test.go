package game

import (
	"testing"
)

func TestNewLocation(t *testing.T) {
	location := newLocation(1, 1)
	assertLocationAttributes(location, 1, 1, t)

	location = newLocation(0, 1)
	assertLocationAttributes(location, 0, 1, t)

	location = newLocation(1, 0)
	assertLocationAttributes(location, 1, 0, t)
}

func assertLocationAttributes(location location, x int, y int, t *testing.T) {
	if location.x != x {
		t.Errorf("location.x should have been %d, was %d", x, location.x)
	}
	if location.y != y {
		t.Errorf("location.y should have been %d, was %d", y, location.y)
	}
}
