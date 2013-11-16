import unittest

from engine import types


class ArmorTest(unittest.TestCase):
    def testBodyArmor(self):
        body_armor = types.new_armor('BodyArmor')
        self.assertEqual(body_armor[types.attack_types['bullet']], 1.0)
        self.assertEqual(body_armor[types.attack_types['cannon']], 4.0)

    def testHeavyMetal(self):
        heavy_metal = types.new_armor('HeavyMetal')
        self.assertEqual(heavy_metal[types.attack_types['bullet']], 0.25)
        self.assertEqual(heavy_metal[types.attack_types['cannon']], 1.0)

    def testWeakMetal(self):
        weak_metal = types.new_armor('WeakMetal')
        self.assertEqual(weak_metal[types.attack_types['bullet']], 0.5)
        self.assertEqual(weak_metal[types.attack_types['cannon']], 2.0)

    def testBodyArmorEqualsBodyArmor(self):
        self.assertEqual(types.new_armor('BodyArmor'), types.new_armor('BodyArmor'))

    def testBodyArmorDoesNotEqualModifiedBodyArmor(self):
        modified = types.new_armor('BodyArmor')
        modified[types.attack_types['bullet']] = 0.0
        self.assertNotEqual(modified, types.new_armor('BodyArmor'))

    def testBodyArmorDoesNotEqualWeakMetal(self):
        self.assertNotEqual(types.new_armor('BodyArmor'), types.new_armor('WeakMetal'))

    def testChangeMapSameValues(self):
        modified = types.new_armor('BodyArmor')
        modified.dictionary = {
            types.attack_types['bullet']: 1.0,
            types.attack_types['cannon']: 4.0}
        self.assertEqual(modified, types.new_armor('BodyArmor'))

    def testChangeMapDifferentValues(self):
        modified = types.new_armor('BodyArmor')
        modified.dictionary = {
            types.attack_types['bullet']: 1.5,
            types.attack_types['cannon']: 4.0}
        self.assertNotEqual(modified, types.new_armor('BodyArmor'))

if __name__ == '__main__':
    unittest.main()
