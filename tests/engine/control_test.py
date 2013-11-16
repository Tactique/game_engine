import unittest

from engine import types, control, consts


class ControlTest(unittest.TestCase):
    def oneVsOther(self, AttackUnit, DefendUnit, delta, start_damage=0):
        attack_unit = types.new_unit(AttackUnit, consts.RED)
        attack_unit.health -= start_damage
        expected_attack_unit = types.new_unit(AttackUnit, consts.RED)
        expected_attack_unit.health -= start_damage
        defend_unit = types.new_unit(DefendUnit, consts.BLUE)
        expected_defend_unit = types.new_unit(DefendUnit, consts.BLUE)
        expected_defend_unit.health -= delta

        control.DoDamage(attack_unit, defend_unit)

        self.assertEqual(attack_unit, expected_attack_unit)
        self.assertEqual(defend_unit, expected_defend_unit)

    # Tank
    def testTankVsTank(self):
        self.oneVsOther('Tank', 'Tank', 5)

    def testTankVsInf(self):
        self.oneVsOther('Tank', 'Infantry', 20)

    def testTankVsRecon(self):
        self.oneVsOther('Tank', 'Recon', 10)

    # Infantry
    def testInfVsTank(self):
        self.oneVsOther('Infantry', 'Tank', 1)

    def testInfVsInf(self):
        self.oneVsOther('Infantry', 'Infantry', 5)

    def testInfVsRecon(self):
        self.oneVsOther('Infantry', 'Recon', 2)

    # Recon
    def testReconVsTank(self):
        self.oneVsOther('Recon', 'Tank', 2)

    def testReconVsInf(self):
        self.oneVsOther('Recon', 'Infantry', 10)

    def testReconVsRecon(self):
        self.oneVsOther('Recon', 'Recon', 5)

    #TODO Add more tests here
    #Damaged
    def testDamageTankVsTank(self):
        self.oneVsOther('Tank', 'Tank', 2, 5)

if __name__ == '__main__':
    unittest.main()
