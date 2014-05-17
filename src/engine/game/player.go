package game

import (
	"api"
)

type Player struct {
	playerId int
	nation   nation
	team     team
}

func NewPlayer(playerId int, nation nation, team team) *Player {
	return &Player{playerId: playerId, nation: nation, team: team}
}

func (player *Player) Serialize() *api.PlayerStruct {
	return &api.PlayerStruct{
		Nation: int(player.nation),
		Team:   int(player.team)}
}

func (player *Player) GetPlayerId() int {
	return player.playerId
}

func (player *Player) GetNation() int {
	return int(player.nation)
}
