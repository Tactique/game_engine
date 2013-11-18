import unittest

from lib import base

from engine import types


class TypesTest(unittest.TestCase):
    def testAttackTypes(self):
        self.assertIsInstance(types.attack_types, base.BaseEnum)
        self.assertGreater(len(types.attack_types), 0)
        for name, value in types.attack_types.items():
            self.assertIsInstance(name, str)
            self.assertIsInstance(value, int)

    def testAttacks(self):
        self.assertIsInstance(types.attacks, list)
        self.assertIsInstance(types.new_attack, type(lambda x: x))
        #TODO Test actually calling function
        #TODO don't do type(lambda) grrrrrr

if __name__ == '__main__':
    unittest.main()
