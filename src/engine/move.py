from lib import contract

from . import movement, loc


@contract.accepts(int, int)
@contract.returns(int)
def get_distance(first, second):
    return abs(first - second)


@contract.accepts([loc.Loc])
@contract.returns(bool)
def assert_all_tiles_touch(move_list):
    old_loc = move_list[0]
    for loc_ in move_list[1:]:
        if (get_distance(loc_.x, old_loc.x) + get_distance(loc_.y, old_loc.y)) != 1:
            return False
        old_loc = loc_
    return True


@contract.accepts(int, movement.Movement, [int], [loc.Loc])
@contract.returns(bool)
def valid_move(moves, movement_type, tiles, move_list):
    if not assert_all_tiles_touch(move_list):
        return False
    for tile in tiles:
        moves -= movement_type[tile]
    return moves > 0
