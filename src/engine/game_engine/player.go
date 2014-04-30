package game_engine

import (
	"api"
)

type player struct {
	playerId int
	nation   nation
	team     team
}

func newPlayer(playerId int, nation nation, team team) *player {
	return &player{playerId: playerId, nation: nation, team: team}
}

func (player *player) serialize() *api.PlayerStruct {
	return &api.PlayerStruct{
		Nation: int(player.nation),
		Team:   int(player.team)}
}
