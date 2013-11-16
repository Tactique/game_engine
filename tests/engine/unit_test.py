import unittest

from engine import consts, types


class UnitTest(unittest.TestCase):
    def testTank(self):
        tank = types.new_unit('Tank', consts.RED)
        self.assertEqual(tank.health, 10)
        self.assertEqual(
            tank.attacks, [types.new_attack('RegularCannon'), types.new_attack('MachineGun')])
        self.assertEqual(tank.armor, types.new_armor('HeavyMetal'))
        self.assertEqual(tank.movement, types.new_movement('Treads'))
        self.assertEqual(tank.distance, 7)

    def testInfantry(self):
        inf = types.new_unit('Infantry', consts.RED)
        self.assertEqual(inf.health, 10)
        self.assertEqual(inf.attacks, [types.new_attack('MachineGun')])
        self.assertEqual(inf.armor, types.new_armor('BodyArmor'))
        self.assertEqual(inf.movement, types.new_movement('Feet'))
        self.assertEqual(inf.distance, 3)

    def testRecon(self):
        recon = types.new_unit('Recon', consts.RED)
        self.assertEqual(recon.health, 10)
        self.assertEqual(recon.attacks, [types.new_attack('DoubleMachineGun')])
        self.assertEqual(recon.armor, types.new_armor('WeakMetal'))
        self.assertEqual(recon.movement, types.new_movement('Tires'))
        self.assertEqual(recon.distance, 9)

    def testTankEqualsTank(self):
        self.assertEqual(types.new_unit('Tank', consts.RED), types.new_unit('Tank', consts.RED))

    def testTankDoesNotEqualRecon(self):
        self.assertNotEqual(
            types.new_unit('Tank', consts.RED),
            types.new_unit('Recon', consts.RED))

    def testTankDoesNotEqualModifiedTank(self):
        modified = types.new_unit('Tank', consts.RED)
        modified.health -= 1
        self.assertNotEqual(modified, types.new_unit('Tank', consts.RED))

        modified = types.new_unit('Tank', consts.RED)
        modified.armor[0] = 1000.0
        self.assertNotEqual(modified, types.new_unit('Tank', consts.RED))

if __name__ == '__main__':
    unittest.main()
