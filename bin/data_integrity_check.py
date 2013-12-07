#!/usr/bin/env python

from engine import file_loader
import notify

notifier = notify.Notifier('Data Integrity')


def verify_all():
    attack_types = verify_enum('attack_types')
    attacks = verify_used('attacks', 'attack_type', attack_types)
    armors = verify_contains('armors', 'resists', attack_types)
    tiles = verify_enum('tiles')
    movements = verify_contains('movements', 'speeds', tiles)
    verify_units(attacks, armors, movements)


def verify_units(attacks, armors, movements):
    def get_name(dict_):
        return dict_['name']

    verify_used('units', 'attacks', map(get_name, attacks))
    verify_used('units', 'armor', map(get_name, armors))
    verify_used('units', 'movement', map(get_name, movements))


def verify_enum(module):
    attack_types = file_loader.read_and_parse_json(module)[0]
    if len(set(attack_types)) != len(attack_types):
        print_failure(module, '', 'Duplicates', attack_types)
    else:
        print_success(module, '', 'No duplicates', attack_types)
    return attack_types


def verify_contains(module, sub_name, sub_set):
    armors = file_loader.read_and_parse_json(module)
    for armor in armors:
        instance = armor['name']
        used = set()
        for resist in armor[sub_name]:
            used.add(resist)

        _verify_exact_subset(module, sub_name, used, sub_set)
    return armors


def verify_used(module, sub_name, sub_set):
    used = set()
    attacks = file_loader.read_and_parse_json(module)
    for sub_attacks in attacks:
        sub_ = sub_attacks[sub_name]
        if not hasattr(sub_, '__iter__'):
            sub_ = (sub_,)
        for attack in sub_:
            used.add(attack)

    _verify_exact_subset(module, sub_name, used, sub_set)

    return attacks


def _verify_exact_subset(module, sub_name, used, sub_set):
    unused = set(sub_set) - used
    unimplemented = used - set(sub_set)
    if len(unused) != 0:
        print_failure(module, sub_name.title(), 'Unused', unused)
    elif len(unimplemented) != 0:
        print_failure(module, sub_name.title(), 'Unimplemented', unimplemented)
    else:
        print_success(module, sub_name.title(), 'All used', used)


def print_failure(module, instance, message, components):
    notifier.failure(_print_message(module, instance, message, components))


def print_success(module, instance, message, components):
    notifier.success(_print_message(module, instance, message, components))


def _print_message(module, instance, message, components, pretty=True):
    return '%s -> %s %s : %s' % (
        notify.trail_spaces(module.title(), 12),
        notify.trail_spaces(instance.title(), 12),
        notify.trail_spaces(message, 14),
        ', '.join(components))


if __name__ == '__main__':
    verify_all()
