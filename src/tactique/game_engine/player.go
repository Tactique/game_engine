package game_engine

import (
    "requests"
)

type player struct {
    playerId int
    nation nation
    team team
}

func newPlayer(playerId int, nation nation, team team) *player {
    return &player{playerId: playerId, nation: nation, team: team}
}

func (player *player) serialize() *requests.PlayerStruct {
    return &requests.PlayerStruct{
        PlayerId: player.playerId,
        Nation: int(player.nation),
        Team: int(player.team)}
}
