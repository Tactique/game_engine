import unittest

from engine import armor, attack


class ArmorTest(unittest.TestCase):
    def testBodyArmor(self):
        body_armor = armor.BodyArmor()
        self.assertEqual(body_armor[attack.BULLET], 1.0)
        self.assertEqual(body_armor[attack.CANNON], 4.0)

    def testHeavyMetal(self):
        heavy_metal = armor.HeavyMetal()
        self.assertEqual(heavy_metal[attack.BULLET], 0.25)
        self.assertEqual(heavy_metal[attack.CANNON], 1.0)

    def testWeakMetal(self):
        weak_metal = armor.WeakMetal()
        self.assertEqual(weak_metal[attack.BULLET], 0.5)
        self.assertEqual(weak_metal[attack.CANNON], 2.0)

    def testBodyArmorEqualsBodyArmor(self):
        self.assertEqual(armor.BodyArmor(), armor.BodyArmor())

    def testBodyArmorDoesNotEqualModifiedBodyArmor(self):
        modified = armor.BodyArmor()
        modified[attack.BULLET] = 0.0
        self.assertNotEqual(modified, armor.BodyArmor())

    def testBodyArmorDoesNotEqualWeakMetal(self):
        self.assertNotEqual(armor.BodyArmor(), armor.WeakMetal())

    def testChangeMapSameValues(self):
        modified = armor.BodyArmor()
        modified.armor = {attack.BULLET: 1.0, attack.CANNON: 4.0}
        self.assertEqual(modified, armor.BodyArmor())

    def testChangeMapDifferentValues(self):
        modified = armor.BodyArmor()
        modified.armor = {attack.BULLET: 1.5, attack.CANNON: 4.0}
        self.assertNotEqual(modified, armor.BodyArmor())

if __name__ == '__main__':
    unittest.main()
