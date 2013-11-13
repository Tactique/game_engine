from lib import contract

from . import tile, base


class Movement(base.BaseDictionary):
    pass


class Treads(Movement):
    def __init__(self):
        self.dictionary = {
            tile.PLAIN: 1.0,
            tile.CITY: 1.0,
            tile.WOODS: 1.0,
            tile.MOUNTAIN: 0.0,
        }


class Tires(Movement):
    def __init__(self):
        self.dictionary = {
            tile.PLAIN: 1.5,
            tile.CITY: 1.0,
            tile.WOODS: 2.0,
            tile.MOUNTAIN: 0.0,
        }


class Feet(Movement):
    def __init__(self):
        self.dictionary = {
            tile.PLAIN: 1.0,
            tile.CITY: 1.0,
            tile.WOODS: 1.0,
            tile.MOUNTAIN: 2.0,
        }
