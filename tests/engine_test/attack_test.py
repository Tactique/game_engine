import unittest

from engine import attack


class AttackTest(unittest.TestCase):
    def testRegularCannon(self):
        regular_cannon = attack.RegularCannon()
        self.assertEqual(regular_cannon.power, 5)
        self.assertEqual(regular_cannon.attackType, attack.CANNON)

    def testMachineGun(self):
        machine_gun = attack.MachineGun()
        self.assertEqual(machine_gun.power, 5)
        self.assertEqual(machine_gun.attackType, attack.BULLET)

    def testDoubleMachineGun(self):
        double_machine_gun = attack.DoubleMachineGun()
        self.assertEqual(double_machine_gun.power, 10)
        self.assertEqual(double_machine_gun.attackType, attack.BULLET)

    def testRegularCannonEqualsRegularCannon(self):
        self.assertEqual(attack.RegularCannon(), attack.RegularCannon())

    def testRegularCannonDoesNotEqualModifiedRegularCannon(self):
        modified = attack.RegularCannon()
        modified.power += 1
        self.assertNotEqual(modified, attack.RegularCannon())

        modified = attack.RegularCannon()
        modified.attackType = -1
        self.assertNotEqual(modified, attack.RegularCannon())

    def testRegularCannonDoesNotEqualMachineGun(self):
        self.assertNotEqual(attack.RegularCannon(), attack.MachineGun())

if __name__ == '__main__':
    unittest.main()
