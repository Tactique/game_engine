import json

from lib import contract

from . import unit, player, types, move, consts


class World(object):
    @contract.self_accepts([int])
    def __init__(self, player_ids):
        self.terrain = [[types.tiles['plain'] for i in range(10)] for j in range(10)]
        self.players = {}
        self.current_unit_id = 0
        for team, player_id in zip(consts.TEAMS, player_ids):
            self.players[team] = player.Player(player_id)
        self.turn_owner = self.players[consts.RED]

    @contract.self_accepts(int)
    @contract.returns(player.Player)
    def get_player(self, player_id):
        for player_ in self.players.values():
            if player_.player_id == player_id:
                return player_
        else:
            raise Exception("Player not in game")

    @contract.self_accepts(int, unit.Unit)
    @contract.returns(None)
    def add_unit(self, player_id, unit_):
        self.get_player(player_id).add_unit(unit_, self.current_unit_id)
        self.current_unit_id += 1

    @contract.self_accepts()
    @contract.returns(str)
    def to_json(self):
        ret_json = {
            'terrain': self.terrain
        }
        for player in self.players.values():
            ret_json.update(
                {
                    player.player_id: player.serialize(True)
                }
            )
        return json.dumps(ret_json)

    @contract.self_accepts(int, int, [(int,)])
    @contract.returns(bool)
    def move_unit(self, player_id, unit_id, move_list):
        def get_tile_from_coord(coord_tuple):
            x, y = coord_tuple
            return self.terrain[x][y]

        player_ = self.get_player(player_id)
        unit_ = player_.get_unit(unit_id)
        tiles = map(get_tile_from_coord, move_list)

        return move.valid_move(
            unit_.distance,
            unit_.movement,
            tiles,
            move_list)
