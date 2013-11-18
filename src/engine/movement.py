from lib import contract, base

from . import file_loader, tile


class Movement(base.BaseDictionary):
    pass


@contract.accepts(tile.Tiles)
def load_movements(tiles):
    args = {}
    names = []
    for movement_ in file_loader.read_and_parse_json('movements'):
        name = str(movement_['name'])
        speeds = {}
        for tile_, multiplier in movement_['speeds'].items():
            speeds[tiles[str(tile_)]] = multiplier
        names.append(name)
        args[name] = speeds

    @contract.accepts(str)
    @contract.returns(Movement)
    def movement_getter(name):
        return Movement(dict(args[name]))
    return movement_getter, names
