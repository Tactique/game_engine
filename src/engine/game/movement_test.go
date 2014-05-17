package game

import (
	"testing"
)

func TestMovement(t *testing.T) {
	movement := newMovement("Treads", 7, map[terrain]multiplier{
		plains: 1})
	assertMovementAttributes(movement, "Treads", 7, map[terrain]multiplier{
		plains: 1}, t)
}

func assertMovementAttributes(movement *movement, name string, distance int, speeds map[terrain]multiplier, t *testing.T) {
	if movement.name != name {
		t.Fatalf("movement.name should have been \"%s\" but was \"%s\"", name, movement.name)
	}
	if movement.distance != distance {
		t.Fatalf("movement.distance should have been %d but was %d", distance, movement.distance)
	}
	if movement.costs[plains] != speeds[plains] {
		t.Fatalf("movement.costs[plains] should have been %d but was %d", speeds[plains], movement.costs[plains])
	}
}
