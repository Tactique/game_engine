from lib import contract

from . import tile


class Movement(object):
    @contract.returns(bool)
    def __eq__(self, movement):
        return (
            type(self) == type(movement) and
            self.movement == movement.movement)

    @contract.returns(bool)
    def __ne__(self, movement):
        return not self.__eq__(movement)

    @contract.self_accepts(int, float)
    @contract.returns(None)
    def __setitem__(self, index, multiplier):
        self.movement[index] = multiplier

    @contract.self_accepts(int)
    @contract.returns(float)
    def __getitem__(self, index):
        return self.movement[index]


class Treads(Movement):
    def __init__(self):
        self.movement = {
            tile.PLAIN: 1.0,
            tile.CITY: 1.0,
            tile.WOODS: 1.0,
            tile.MOUNTAIN: 0.0,
        }


class Tires(Movement):
    def __init__(self):
        self.movement = {
            tile.PLAIN: 1.5,
            tile.CITY: 1.0,
            tile.WOODS: 2.0,
            tile.MOUNTAIN: 0.0,
        }


class Feet(Movement):
    def __init__(self):
        self.movement = {
            tile.PLAIN: 1.0,
            tile.CITY: 1.0,
            tile.WOODS: 1.0,
            tile.MOUNTAIN: 2.0,
        }
