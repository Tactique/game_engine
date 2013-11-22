#!/usr/bin/env python

from engine import file_loader


def verbose_print(line):
    print 'OK: %s' % (line,)


def verify_enum(dir_name, name):
    attack_types = file_loader.read_and_parse_json(dir_name)[0]
    if len(set(attack_types)) != len(attack_types):
        raise Exception('%s | Duplicate(s) found : %s' % (name, ', '.join(attack_types),))
    else:
        verbose_print('%s | No duplicates found in : %s' % (name, ', '.join(attack_types),))
    return attack_types


def verify_contains(dir_name, module_name, sub_name, sub_set):
    armors = file_loader.read_and_parse_json(dir_name)
    for armor in armors:
        name = armor['name']
        used_resists = set()
        for resist in armor[sub_name]:
            used_resists.add(resist)
        unused_resists = set(sub_set) - used_resists
        unimplemented_resists = used_resists - set(sub_set)
        if len(unused_resists) != 0:
            raise Exception('%s | %s : Unused %s : %s' % (
                module_name, name, sub_name, ', '.join(unused_resists),))
        if len(unimplemented_resists) != 0:
            raise Exception('%s | %s : Unimplemented %s : %s' % (
                module_name, name, sub_name, ', '.join(unimplemented_resists),))
        else:
            verbose_print('%s | %s : implements all of %s' % (
                module_name, name, ', '.join(used_resists)))
    return armors


def verify_used(dir_name, module_name, sub_name, sub_set):
    used_attack_types = set()
    attacks = file_loader.read_and_parse_json(dir_name)
    for sub_attacks in attacks:
        sub_ = sub_attacks[sub_name]
        if not hasattr(sub_, '__iter__'):
            sub_ = (sub_,)
        for attack in sub_:
            used_attack_types.add(attack)

    unused_attack_types = set(sub_set) - used_attack_types
    unimplemented_attack_types = used_attack_types - set(sub_set)
    if len(unused_attack_types) != 0:
        raise Exception('%s | Unused %s: %s' % (
            module_name, sub_name, ', '.join(unused_attack_types)))
    elif len(unimplemented_attack_types) != 0:
        raise Exception('%s | Unimplemented %s: %s' % (
            module_name, sub_name, ', '.join(unimplemented_attack_types)))
    else:
        verbose_print('%s | All %s used : %s' % (
            module_name, sub_name, ', '.join(sub_set),))
    return attacks


def verify_attack_types():
    return verify_enum('attack_types', 'Attack Type')


def verify_attacks(attack_types):
    return verify_used('attacks', 'Attack', 'attack_type', attack_types)


def verify_armors(attack_types):
    return verify_contains('armors', 'Armor', 'resists', attack_types)


def verify_tiles():
    return verify_enum('tiles', 'Tiles')


def verify_movements(tiles):
    return verify_contains('movements', 'Movement', 'speeds', tiles)


def verify_units(attacks, armors, movements):
    def get_name(dict_):
        return dict_['name']

    verify_used('units', 'Unit', 'attacks', map(get_name, attacks))
    verify_used('units', 'Unit', 'armor', map(get_name, armors))
    verify_used('units', 'Unit', 'movement', map(get_name, movements))


def verify_all():
    attack_types = verify_attack_types()
    attacks = verify_attacks(attack_types)
    armors = verify_armors(attack_types)
    tiles = verify_tiles()
    movements = verify_movements(tiles)
    units = verify_units(attacks, armors, movements)


if __name__ == '__main__':
    verify_all()
