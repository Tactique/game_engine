from lib import contract

from . import movement


@contract.accepts(int, int)
@contract.returns(int)
def get_distance(first, second):
    return abs(first - second)


@contract.accepts(list)
@contract.returns(bool)
def assert_all_tiles_touch(move_list):
    old_x, old_y = move_list[0]
    for x, y in move_list[1:]:
        if (get_distance(x, old_x) + get_distance(y, old_y)) != 1:
            return False
        old_x, old_y = x, y
    return True


@contract.accepts(int, movement.Movement, list, list)
@contract.returns(bool)
def valid_move(moves, movement_type, tiles, move_list):
    if not assert_all_tiles_touch(move_list):
        return False
    for tile in tiles:
        moves -= movement_type[tile]
    return moves > 0
