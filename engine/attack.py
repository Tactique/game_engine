from lib import contract

BULLET = 0
CANNON = 1

from . import base


class Attack(base.BaseClass):
    @contract.self_accepts(int, int)
    def __init__(self, power, attackType):
        self.power = power
        self.attackType = attackType

    @contract.returns(bool)
    def __eq__(self, attack):
        return (
            type(self) == type(attack) and
            self.power == attack.power and
            self.attackType == attack.attackType)


class RegularCannon(Attack):
    def __init__(self):
        Attack.__init__(self, 5, CANNON)


class MachineGun(Attack):
    def __init__(self):
        Attack.__init__(self, 5, BULLET)


class DoubleMachineGun(Attack):
    def __init__(self):
        Attack.__init__(self, 10, BULLET)
