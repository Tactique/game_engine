package game_engine

import (
    "testing"
)

func TestNewGame(t *testing.T) {
    world, err := NewGame([]int{26, 13}, 1)
    if err != nil {
        t.Fatalf("Problem building world %s", err)
    }
    if world.numPlayers != 2 {
        t.Fatalf("Wrong number of players (field), should be 2, was %d", world.numPlayers)
    }
    if len(world.players) != 2 {
        t.Fatalf("Wrong number of players (len), should be 2, was %d", len(world.players))
    }
    assertPlayerAttributes(world.players[0], 26, 0, 0, t)
    assertPlayerAttributes(world.players[1], 13, 1, 1, t)
}
