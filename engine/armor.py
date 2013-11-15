from lib import contract

from . import types, base


class Armor(base.BaseDictionary):
    pass


class BodyArmor(Armor):
    def __init__(self):
        self.dictionary = {
            types.attack_types['bullet']: 1.0,
            types.attack_types['cannon']: 4.0,
        }


class HeavyMetal(Armor):
    def __init__(self):
        self.dictionary = {
            types.attack_types['bullet']: 0.25,
            types.attack_types['cannon']: 1.0,
        }


class WeakMetal(Armor):
    def __init__(self):
        self.dictionary = {
            types.attack_types['bullet']: 0.5,
            types.attack_types['cannon']: 2.0,
        }
