import unittest

from engine import unit, attack, armor, movement, consts


class UnitTest(unittest.TestCase):
    def testTank(self):
        tank = unit.Tank(consts.RED)
        self.assertEqual(tank.health, 10)
        self.assertEqual(
            tank.attacks, [attack.RegularCannon(), attack.MachineGun()])
        self.assertEqual(tank.armor, armor.HeavyMetal())
        self.assertEqual(tank.movement, movement.Treads())
        self.assertEqual(tank.distance, 7)

    def testInfantry(self):
        inf = unit.Infantry(consts.RED)
        self.assertEqual(inf.health, 10)
        self.assertEqual(inf.attacks, [attack.MachineGun()])
        self.assertEqual(inf.armor, armor.BodyArmor())
        self.assertEqual(inf.movement, movement.Feet())
        self.assertEqual(inf.distance, 3)

    def testRecon(self):
        recon = unit.Recon(consts.RED)
        self.assertEqual(recon.health, 10)
        self.assertEqual(recon.attacks, [attack.DoubleMachineGun()])
        self.assertEqual(recon.armor, armor.WeakMetal())
        self.assertEqual(recon.movement, movement.Tires())
        self.assertEqual(recon.distance, 9)

    def testUnitToString(self):
        unit_repr = unit.Tank(consts.RED).toString()

    def testTankEqualsTank(self):
        self.assertEqual(unit.Tank(consts.RED), unit.Tank(consts.RED))

    def testTankDoesNotEqualRecon(self):
        self.assertNotEqual(unit.Tank(consts.RED), unit.Recon(consts.RED))

    def testTankDoesNotEqualModifiedTank(self):
        modified = unit.Tank(consts.RED)
        modified.health -= 1
        self.assertNotEqual(modified, unit.Recon(consts.RED))

if __name__ == '__main__':
    unittest.main()
