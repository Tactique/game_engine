from lib import contract

from . import unit


class Player(object):
    @contract.self_accepts(int)
    def __init__(self, uid):
        self.units = []
        self.uid = uid

    @contract.self_accepts(unit.Unit)
    @contract.returns(None)
    def add_unit(self, unit_):
        self.units.append(unit_)
