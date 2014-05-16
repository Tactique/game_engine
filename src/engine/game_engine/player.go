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

func (player *player) Serialize() *api.PlayerStruct {
	return &api.PlayerStruct{
		Nation: int(player.nation),
		Team:   int(player.team)}
}

func (player *player) GetPlayerId() int {
	return player.playerId
}

func (player *player) GetNation() int {
	return int(player.nation)
}
