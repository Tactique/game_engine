from lib import contract, base

from . import file_loader


class Tiles(base.BaseEnum):
    pass


@contract.returns(Tiles)
def load_tiles():
    tiles = Tiles()
    tile_list = file_loader.read_and_parse_json('tiles')[0]
    for enumeration, tile_ in enumerate(tile_list):
        tiles[str(tile_)] = enumeration
    return tiles
