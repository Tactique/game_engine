from lib import contract

from . import attack, base


class Armor(base.BaseDictionary):
    pass


class BodyArmor(Armor):
    def __init__(self):
        self.dictionary = {
            attack.BULLET: 1.0,
            attack.CANNON: 4.0,
        }


class HeavyMetal(Armor):
    def __init__(self):
        self.dictionary = {
            attack.BULLET: 0.25,
            attack.CANNON: 1.0,
        }


class WeakMetal(Armor):
    def __init__(self):
        self.dictionary = {
            attack.BULLET: 0.5,
            attack.CANNON: 2.0,
        }
