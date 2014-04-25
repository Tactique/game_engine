package game_engine

import (
    "testing"
)

func TestNewPlayer(t *testing.T) {
    player := newPlayer(26, 0, 0)
    assertPlayerAttributes(player, 26, 0, 0, t)

    player = newPlayer(13, 1, 1)
    assertPlayerAttributes(player, 13, 1, 1, t)

    player = newPlayer(10, 1, 2)
    assertPlayerAttributes(player, 10, 1, 2, t)
}

func assertPlayerAttributes(player *player, playerId int, nation nation,  team team, t *testing.T) {
    if player.playerId != playerId {
        t.Fatalf("player.playerId should be %d was %d", playerId, player.playerId)
    }
    if player.nation != nation {
        t.Fatalf("player.nation should be %d was %d", nation, player.nation)
    }
    if player.team != team {
        t.Fatalf("player.team should be %d was %d", team, player.team)
    }
}
