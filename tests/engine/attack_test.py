import unittest

from engine import types


class AttackTest(unittest.TestCase):
    def testRegularCannon(self):
        regular_cannon = types.new_attack('RegularCannon')
        self.assertEqual(regular_cannon.power, 5)
        self.assertEqual(regular_cannon.attackType, types.attack_types['cannon'])

    def testMachineGun(self):
        machine_gun = types.new_attack('MachineGun')
        self.assertEqual(machine_gun.power, 5)
        self.assertEqual(machine_gun.attackType, types.attack_types['bullet'])

    def testDoubleMachineGun(self):
        double_machine_gun = types.new_attack('DoubleMachineGun')
        self.assertEqual(double_machine_gun.power, 10)
        self.assertEqual(double_machine_gun.attackType, types.attack_types['bullet'])

    def testRegularCannonEqualsRegularCannon(self):
        self.assertEqual(types.new_attack('RegularCannon'), types.new_attack('RegularCannon'))

    def testRegularCannonDoesNotEqualModifiedRegularCannon(self):
        modified = types.new_attack('RegularCannon')
        modified.power += 1
        self.assertNotEqual(modified, types.new_attack('RegularCannon'))

        modified = types.new_attack('RegularCannon')
        modified.attackType = -1
        self.assertNotEqual(modified, types.new_attack('RegularCannon'))

    def testRegularCannonDoesNotEqualMachineGun(self):
        self.assertNotEqual(types.new_attack('RegularCannon'), types.new_attack('MachineGun'))

if __name__ == '__main__':
    unittest.main()
