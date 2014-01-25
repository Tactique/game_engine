from lib import contract, base

from . import file_loader, tile


class Movement(base.BaseDictionary):
    def __init__(self, name, dictionary):
        base.BaseDictionary.__init__(self, dictionary)
        self.name = name

    @contract.self_accepts(bool)
    @contract.returns(dict)
    def serialize(self, public):
        return {
            "speeds": self.dictionary,
            "name": self.name}


@contract.accepts(tile.Tiles)
def load_movements(tiles):
    @contract.accepts(str)
    @contract.returns(Movement)
    def movement_getter(name):
        movement_ = args[name]
        speeds = {}
        movement_name = movement_['name']
        for tile_, multiplier in movement_['speeds'].items():
            speeds[tiles[str(tile_)]] = multiplier
        return Movement(movement_name, speeds)

    args = file_loader.load_struct('movements')
    return movement_getter, args.keys()
