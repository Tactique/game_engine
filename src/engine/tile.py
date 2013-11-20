from lib import contract, base

from . import file_loader


class Tiles(base.BaseEnum):
    pass


@contract.returns(Tiles)
def load_tiles():
    return Tiles(file_loader.load_enum('tiles'))
