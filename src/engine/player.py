from lib import contract

from . import unit, consts


class Player(object):
    @contract.self_accepts(int, consts.Nation)
    def __init__(self, player_id, nation_):
        self.player_id = player_id
        self.nation = nation_
        self.team = consts.Team(nation_)

    @contract.self_accepts(bool)
    def serialize(self, public):
        return {
            "player_id": self.player_id,
            "nation": self.nation.serialize(public),
            "team": self.team.serialize(public)
        }
