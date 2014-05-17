package game

import (
	"testing"
)

func TestNewUnit(t *testing.T) {
	movement := newMovement("Treads", 7, map[terrain]multiplier{
		plains: 1})
	unit := newUnit(
		"Tank", red, movement)
	assertUnitAttributes(unit, "Tank", red, movement, t)
}

func assertUnitAttributes(unit *unit, name string, nation nation, movement *movement, t *testing.T) {
	if unit.name != name {
		t.Fatalf("unit.name should be \"%s\" got \"%s\"", name, unit.name)
	}
	if unit.nation != nation {
		t.Fatalf("unit.nation should be %d got %d", nation, unit.nation)
	}
	if unit.health != 10 {
		t.Fatalf("unit.health should be 10 got %s", unit.health)
	}
	if unit.movement != movement {
		t.Fatalf("unit.movement should be %s got %s", movement.name, unit.movement.name)
	}
}
