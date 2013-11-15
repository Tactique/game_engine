from lib import contract

from . import consts, types, armor, movement, base, file_loader


class Unit(base.BaseClass):
    @contract.self_accepts(list, armor.Armor, movement.Movement, int, consts.Team)
    def __init__(self, attacks_, armor_, movement_, distance_, team_):
        self.health = consts.MAX_HEALTH
        self.attacks = attacks_
        self.armor = armor_
        self.movement = movement_
        self.distance = distance_
        self.team = team_


class Tank(Unit):
    def __init__(self, team_):
        Unit.__init__(
            self,
            [types.new_attack('RegularCannon'), types.new_attack('MachineGun')],
            armor.HeavyMetal(),
            movement.Treads(),
            7,
            team_)


class Infantry(Unit):
    def __init__(self, team_):
        Unit.__init__(
            self,
            [types.new_attack('MachineGun')],
            armor.BodyArmor(),
            movement.Feet(),
            3,
            team_)


class Recon(Unit):
    def __init__(self, team_):
        Unit.__init__(
            self,
            [types.new_attack('DoubleMachineGun')],
            armor.WeakMetal(),
            movement.Tires(),
            9,
            team_)
