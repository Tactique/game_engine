from lib import contract

from . import consts, attack, armor, movement, base


class Unit(base.BaseClass):
    @contract.self_accepts(list, armor.Armor, movement.Movement, int, consts.Team)
    def __init__(self, attacks_, armor_, movement_, distance_, team_):
        self.health = consts.MAX_HEALTH
        self.attacks = attacks_
        self.armor = armor_
        self.movement = movement_
        self.distance = distance_
        self.team = team_

    @contract.returns(bool)
    def __eq__(self, unit):
        return (
            type(self) == type(unit) and
            self.health == unit.health and
            self.attacks == unit.attacks and
            self.armor == unit.armor and
            self.movement == unit.movement and
            self.distance == unit.distance and
            self.team == self.team)

    def toString(self):
        return 'Health : %s\nAttacks : %s\nArmor : %s\nMovement : %s' % (
            self.health, self.attacks, self.armor, self.movement)


class Tank(Unit):
    def __init__(self, team_):
        Unit.__init__(
            self,
            [attack.RegularCannon(), attack.MachineGun()],
            armor.HeavyMetal(),
            movement.Treads(),
            7,
            team_)


class Infantry(Unit):
    def __init__(self, team_):
        Unit.__init__(
            self,
            [attack.MachineGun()],
            armor.BodyArmor(),
            movement.Feet(),
            3,
            team_)


class Recon(Unit):
    def __init__(self, team_):
        Unit.__init__(
            self,
            [attack.DoubleMachineGun()],
            armor.WeakMetal(),
            movement.Tires(),
            9,
            team_)
