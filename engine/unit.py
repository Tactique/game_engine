from lib import contract

from . import consts, attack, armor, movement


class Unit(object):
    @contract.self_accepts(list, armor.Armor, movement.Movement)
    def __init__(self, attacks, armor, movement):
        self.health = consts.MAX_HEALTH
        self.attacks = attacks
        self.armor = armor
        self.movement = movement

    def __eq__(self, unit):
        return (
            type(self) == type(unit) and
            self.health == unit.health and
            self.attacks == unit.attacks and
            self.armor == unit.armor and
            self.movement == unit.movement)

    @contract.returns(bool)
    def __ne__(self, unit):
        return not self.__eq__(unit)

    def toString(self):
        return 'Health : %s\nAttacks : %s\nArmor : %s\nMovement : %s' % (
            self.health, self.attacks, self.armor, self.movement)


class Tank(Unit):
    def __init__(self):
        Unit.__init__(
            self,
            [attack.RegularCannon(), attack.MachineGun()],
            armor.HeavyMetal(),
            movement.Treads())


class Infantry(Unit):
    def __init__(self):
        Unit.__init__(
            self,
            [attack.MachineGun()],
            armor.BodyArmor(),
            movement.Feet())


class Recon(Unit):
    def __init__(self):
        Unit.__init__(
            self,
            [attack.DoubleMachineGun()],
            armor.WeakMetal(),
            movement.Tires())
