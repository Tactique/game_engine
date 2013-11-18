from lib import contract

from . import file_loader


@contract.returns(dict)
def load_tiles():
    tiles = {}
    tile_list = file_loader.read_and_parse_json('tiles')[0]
    for enumeration, tile_ in enumerate(tile_list):
        tiles[str(tile_)] = enumeration
    return tiles
