import json

from lib import contract

from . import unit, player, tile


class World(object):
    @contract.self_accepts(list)
    def __init__(self, uids):
        self.terrain = [[tile.PLAIN for i in range(10)] for i in range(10)]
        self.players = {}
        colors = ['Red', 'Blue', 'Green', 'Yellow']
        for uid, color in zip(uids, colors):
            self.players[color] = player.Player(uid)

    @contract.self_accepts(int)
    @contract.returns(player.Player)
    def get_player(self, uid):
        for player_ in self.players.values():
            if player_.uid == uid:
                return player_
        raise Exception("Player not in game")

    @contract.self_accepts(int, unit.Unit)
    @contract.returns(None)
    def add_unit(self, uid, unit_):
        self.get_player(uid).add_unit(unit_)

    @contract.self_accepts()
    @contract.returns(str)
    def to_json(self):
        return json.dumps(self.terrain)

    @contract.self_accepts(int, tuple, list)
    @contract.returns(bool)
    def move(self, player_id, coord, move_list):
        return False
