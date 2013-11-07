import unittest

from engine import unit
from engine import attack
from engine import armor
from engine import movement


class UnitTest(unittest.TestCase):
    def testTank(self):
        tank = unit.Tank()
        self.assertEqual(tank.health, 10)
        self.assertEqual(
            tank.attacks, [attack.RegularCannon(), attack.MachineGun()])
        self.assertEqual(tank.armor, armor.HeavyMetal())
        self.assertEqual(tank.movement, movement.Treads())

    def testInfantry(self):
        inf = unit.Infantry()
        self.assertEqual(inf.health, 10)
        self.assertEqual(inf.attacks, [attack.MachineGun()])
        self.assertEqual(inf.armor, armor.BodyArmor())
        self.assertEqual(inf.movement, movement.Feet())

    def testRecon(self):
        recon = unit.Recon()
        self.assertEqual(recon.health, 10)
        self.assertEqual(recon.attacks, [attack.DoubleMachineGun()])
        self.assertEqual(recon.armor, armor.WeakMetal())
        self.assertEqual(recon.movement, movement.Tires())

    def testUnitToString(self):
        unit_repr = unit.Tank().toString()

    def testTankEqualsTank(self):
        self.assertEqual(unit.Tank(), unit.Tank())

    def testTankDoesNotEqualRecon(self):
        self.assertNotEqual(unit.Tank(), unit.Recon())

    def testTankDoesNotEqualModifiedTank(self):
        modified = unit.Tank()
        modified.health -= 1
        self.assertNotEqual(modified, unit.Recon())

if __name__ == '__main__':
    unittest.main()
