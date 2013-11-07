from lib import contract

from . import attack


class Armor(object):
    @contract.returns(bool)
    def __eq__(self, armor):
        return self.armor == armor.armor

    @contract.returns(bool)
    def __ne__(self, armor):
        return not self.__eq__(armor)

    @contract.self_accepts(int, float)
    @contract.returns(None)
    def __setitem__(self, index, multiplier):
        self.armor[index] = multiplier

    @contract.self_accepts(int)
    @contract.returns(float)
    def __getitem__(self, index):
        return self.armor[index]


class BodyArmor(Armor):
    def __init__(self):
        self.armor = {
            attack.BULLET: 1.0,
            attack.CANNON: 4.0,
        }


class HeavyMetal(Armor):
    def __init__(self):
        self.armor = {
            attack.BULLET: 0.25,
            attack.CANNON: 1.0,
        }


class WeakMetal(Armor):
    def __init__(self):
        self.armor = {
            attack.BULLET: 0.5,
            attack.CANNON: 2.0,
        }
