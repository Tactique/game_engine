from lib import contract, base

from . import file_loader


class Loc:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def serialize(self, public):
        return {
            "x": self.x,
            "y": self.y,
        }
