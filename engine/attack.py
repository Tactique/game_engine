from lib import contract

BULLET = 0
CANNON = 1


class Attack(object):
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

    @contract.returns(bool)
    def __ne__(self, attack):
        return not self.__eq__(attack)


class RegularCannon(Attack):
    def __init__(self):
        Attack.__init__(self, 5, CANNON)


class MachineGun(Attack):
    def __init__(self):
        Attack.__init__(self, 5, BULLET)


class DoubleMachineGun(Attack):
    def __init__(self):
        Attack.__init__(self, 10, BULLET)
