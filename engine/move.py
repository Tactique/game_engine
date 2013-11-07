from lib import contract

from . import movement


@contract.accepts(int, movement.Movement, list)
@contract.returns(bool)
def valid_move(moves, movement_type, tiles):
    for tile in tiles:
        moves -= movement_type[tile]
    return moves > 0
