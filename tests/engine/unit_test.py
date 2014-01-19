import unittest

from engine import consts, types, loc


class UnitTest(unittest.TestCase):
    def testTank(self):
        tank = types.new_unit('Tank', consts.RED, loc.Loc(1, 1))
        self.assertEqual(tank.health, 10)
        self.assertEqual(
            tank.attacks, [types.new_attack('RegularCannon'), types.new_attack('MachineGun')])
        self.assertEqual(tank.armor, types.new_armor('HeavyMetal'))
        self.assertEqual(tank.movement, types.new_movement('Treads'))
        self.assertEqual(tank.distance, 7)

    def testInfantry(self):
        inf = types.new_unit('Infantry', consts.RED, loc.Loc(1, 1))
        self.assertEqual(inf.health, 10)
        self.assertEqual(inf.attacks, [types.new_attack('MachineGun')])
        self.assertEqual(inf.armor, types.new_armor('BodyArmor'))
        self.assertEqual(inf.movement, types.new_movement('Feet'))
        self.assertEqual(inf.distance, 3)

    def testRecon(self):
        recon = types.new_unit('Recon', consts.RED, loc.Loc(1, 1))
        self.assertEqual(recon.health, 10)
        self.assertEqual(recon.attacks, [types.new_attack('DoubleMachineGun')])
        self.assertEqual(recon.armor, types.new_armor('WeakMetal'))
        self.assertEqual(recon.movement, types.new_movement('Tires'))
        self.assertEqual(recon.distance, 9)

    def testTankEqualsTank(self):
        self.assertEqual(
            types.new_unit('Tank', consts.RED, loc.Loc(1, 1)),
            types.new_unit('Tank', consts.RED, loc.Loc(1, 1)))

    def testTankDoesNotEqualRecon(self):
        self.assertNotEqual(
            types.new_unit('Tank', consts.RED, loc.Loc(1, 1)),
            types.new_unit('Recon', consts.RED, loc.Loc(1, 1)))

    def testTankDoesNotEqualModifiedTank(self):
        modified = types.new_unit('Tank', consts.RED, loc.Loc(1, 1))
        modified.health -= 1
        self.assertNotEqual(modified, types.new_unit('Tank', consts.RED, loc.Loc(1, 1)))

        modified = types.new_unit('Tank', consts.RED, loc.Loc(1, 1))
        modified.armor[0] = 1000.0
        self.assertNotEqual(modified, types.new_unit('Tank', consts.RED, loc.Loc(1, 1)))

    def testSerializeTank(self):
        tank = types.new_unit('Tank', consts.RED, loc.Loc(1, 1))
        self.assertEqual(
            tank.serialize(True),
            {
                'team': 0,
                'name': 'Tank',
                'health': 10,
                #'attacks': self.attacks.serialize(public),
                #'armor': self.armor.serialize(public),
                'movement': {0: 1.0, 1: 1.0, 2: 1.0, 3: 0.0},
                'distance': 7,
                'loc': {'x': 1, 'y': 1},
            }
        )

if __name__ == '__main__':
    unittest.main()
