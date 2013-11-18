from lib import contract, base

from . import consts, armor, movement, file_loader


class Unit(base.BaseClass):
    @contract.self_accepts(consts.Team, str, list, armor.Armor, movement.Movement, int)
    def __init__(self, team_, name, attacks_, armor_, movement_, distance_):
        self.team = team_
        self.name = name
        self.health = consts.MAX_HEALTH
        self.attacks = attacks_
        self.armor = armor_
        self.movement = movement_
        self.distance = distance_


def load_units(new_attack, new_armor, new_movement):
    args = {}
    for unit_ in file_loader.read_and_parse_json('units'):
        name = str(unit_['name'])
        args[name] = unit_

    @contract.accepts(str, consts.Team)
    def unit_getter(name, team_):
        unit_ = args[name]
        attacks = []
        for attack_ in unit_['attacks']:
            attacks.append(new_attack(str(attack_)))
        armor_ = new_armor(str(unit_['armor']))
        movement_ = new_movement(str(unit_['movement']))
        distance = unit_['distance']
        return Unit(team_, name, attacks, armor_, movement_, distance)
    return unit_getter, args.keys()
