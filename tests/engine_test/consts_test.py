import unittest

from engine import consts


class ConstsTest(unittest.TestCase):
    def testMaxHealth(self):
        self.assertEqual(consts.MAX_HEALTH, 10)

    def testTeams(self):
        self.assertEqual(consts.RED.team_id, 0)
        self.assertEqual(consts.BLUE.team_id, 1)
        self.assertEqual(consts.GREEN.team_id, 2)
        self.assertEqual(consts.YELLOW.team_id, 3)

if __name__ == '__main__':
    unittest.main()
