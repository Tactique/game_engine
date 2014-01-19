from lib import contract, base

from . import consts, attack, armor, movement, file_loader, loc


class Unit(base.BaseClass):
    @contract.self_accepts(
        consts.Team, str, [attack.Attack], armor.Armor, movement.Movement, int, loc.Loc)
    def __init__(self, team_, name, attacks_, armor_, movement_, distance_, loc):
        self.team = team_
        self.name = name
        self.health = consts.MAX_HEALTH
        self.attacks = attacks_
        self.armor = armor_
        self.movement = movement_
        self.distance = distance_
        self.loc = loc

    #TODO flesh out all stats
    @contract.self_accepts(bool)
    def serialize(self, public):
        return {
            'team': self.team.serialize(public),
            'name': self.name,
            'health': self.health,
            #'attacks': self.attacks.serialize(public),
            #'armor': self.armor.serialize(public),
            'movement': self.movement.serialize(public),
            'distance': self.distance,
            'loc': self.loc.serialize(public),
        }


def load_units(new_attack, new_armor, new_movement):
    @contract.accepts(str, consts.Team, loc.Loc)
    @contract.returns(Unit)
    def unit_getter(name, team_, loc):
        unit_ = args[name]
        attacks = []
        for attack_ in unit_['attacks']:
            attacks.append(new_attack(str(attack_)))
        armor_ = new_armor(str(unit_['armor']))
        movement_ = new_movement(str(unit_['movement']))
        distance = unit_['distance']
        return Unit(team_, name, attacks, armor_, movement_, distance, loc)

    args = file_loader.load_struct('units')
    return unit_getter, args.keys()
