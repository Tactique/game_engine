import unittest

from engine import unit, control, consts


class ControlTest(unittest.TestCase):
    def oneVsOther(self, AttackUnit, DefendUnit, delta, start_damage=0):
        attack_unit = AttackUnit(consts.RED)
        attack_unit.health -= start_damage
        expected_attack_unit = AttackUnit(consts.RED)
        expected_attack_unit.health -= start_damage
        defend_unit = DefendUnit(consts.BLUE)
        expected_defend_unit = DefendUnit(consts.BLUE)
        expected_defend_unit.health -= delta

        control.DoDamage(attack_unit, defend_unit)

        self.assertEqual(attack_unit, expected_attack_unit)
        self.assertEqual(defend_unit, expected_defend_unit)

    # Tank
    def testTankVsTank(self):
        self.oneVsOther(unit.Tank, unit.Tank, 5)

    def testTankVsInf(self):
        self.oneVsOther(unit.Tank, unit.Infantry, 20)

    def testTankVsRecon(self):
        self.oneVsOther(unit.Tank, unit.Recon, 10)

    # Infantry
    def testInfVsTank(self):
        self.oneVsOther(unit.Infantry, unit.Tank, 1)

    def testInfVsInf(self):
        self.oneVsOther(unit.Infantry, unit.Infantry, 5)

    def testInfVsRecon(self):
        self.oneVsOther(unit.Infantry, unit.Recon, 2)

    # Recon
    def testReconVsTank(self):
        self.oneVsOther(unit.Recon, unit.Tank, 2)

    def testReconVsInf(self):
        self.oneVsOther(unit.Recon, unit.Infantry, 10)

    def testReconVsRecon(self):
        self.oneVsOther(unit.Recon, unit.Recon, 5)

    #TODO Add more tests here
    #Damaged
    def testDamageTankVsTank(self):
        self.oneVsOther(unit.Tank, unit.Tank, 2, 5)

if __name__ == '__main__':
    unittest.main()
