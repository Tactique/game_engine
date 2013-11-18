from . import attack, armor, tile, movement, unit

attack_types = attack.load_attack_types()

new_attack, attacks = attack.load_attacks(attack_types)

new_armor, armors = armor.load_armors(attack_types)

tiles = tile.load_tiles()

new_movement, movements = movement.load_movements(tiles)

new_unit, units = unit.load_units(new_attack, new_armor,  new_movement)
