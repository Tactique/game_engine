import json

from lib import contract

from . import unit, player, types, move, consts, loc


class World(object):
    @contract.self_accepts([int])
    def __init__(self, player_ids):
        self.terrain = [[types.tiles['plain'] for i in range(10)] for j in range(10)]
        self.units = {}
        self.players = []
        for nation, player_id in zip(consts.NATIONS, player_ids):
            self.players.append(player.Player(player_id, nation))
        #TODO Random initial player
        self.turn_owner = self.players[0]

    @contract.self_accepts(int)
    @contract.returns(player.Player)
    def get_player(self, player_id):
        for player_ in self.players:
            if player_.player_id == player_id:
                return player_
        else:
            raise Exception("Player not in game")

    @contract.self_accepts(unit.Unit, loc.Loc)
    @contract.returns(None)
    def add_unit(self, unit_, loc_):
        if self.is_valid_placement(loc_):
            self.units[(loc_.x, loc_.y)] = unit_

    def get_unit(self, loc_):
        return self.units[(loc_.x, loc_.y)]

    def is_valid_placement(self, loc_):
        if len(self.terrain) >= loc_.x:
            if len(self.terrain[loc_.x]) >= loc_.y:
                if not (loc_.x, loc_.y) in self.units.keys():
                    return True
        return False

    @contract.self_accepts()
    @contract.returns(str)
    def to_json(self):
        ret_json = {
            'terrain': self.terrain
        }
        players = []
        for player_ in self.players:
            players.append({"player": player_.serialize(True)})
        ret_json.update({"players": players})
        units = []
        for location, unit_ in self.units.items():
            units.append({"unit": unit_.serialize(True, loc.Loc(*location))})
        ret_json.update({"units": units})
        return json.dumps({"world": ret_json})

    @contract.self_accepts([loc.Loc])
    @contract.returns(bool)
    def move_unit(self, move_list):
        def get_tile_from_coord(location):
            return self.terrain[location.x][location.y]

        if len(move_list) < 2:
            print 'move:must move 2'
            return False
        if not (move_list[0].x, move_list[0].y) in self.units.keys():
            print 'not a valid starting point'
            return False
        unit_ = self.get_unit(move_list[0])
        tiles = map(get_tile_from_coord, move_list)

        if move.valid_move(
                unit_.distance,
                unit_.movement,
                tiles,
                move_list):
            unit_ = self.units[(move_list[0].x, move_list[0].y)]
            self.units[(move_list[-1].x, move_list[-1].y)] = unit_
            del self.units[(move_list[0].x, move_list[0].y)]
            return True
        else:
            return False
