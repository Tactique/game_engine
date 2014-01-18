from lib import contract

from . import unit


class Player(object):
    @contract.self_accepts(int)
    def __init__(self, player_id):
        self.units = {}
        self.player_id = player_id

    @contract.self_accepts(unit.Unit, int)
    @contract.returns(None)
    def add_unit(self, unit_, unit_id):
        self.units[unit_id] = unit_

    @contract.self_accepts(int)
    @contract.returns(unit.Unit)
    def get_unit(self, unit_id):
        if unit_id in self.units:
            return self.units[unit_id]
        else:
            raise Exception("Player does not have that unit id")

    @contract.self_accepts(bool)
    def serialize(self, public):
        return {key: val.serialize(public) for key, val in self.units.items()}
