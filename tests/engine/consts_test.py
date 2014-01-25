import unittest

from engine import consts


class ConstsTest(unittest.TestCase):
    def testMaxHealth(self):
        self.assertEqual(consts.MAX_HEALTH, 10)

    def testNations(self):
        self.assertEqual(consts.RED.nation_id, 0)
        self.assertEqual(consts.BLUE.nation_id, 1)
        self.assertEqual(consts.GREEN.nation_id, 2)
        self.assertEqual(consts.YELLOW.nation_id, 3)
        self.assertEqual(
            consts.NATIONS,
            [consts.RED, consts.BLUE, consts.GREEN, consts.YELLOW])

    def testSerializeNations(self):
        self.assertEqual(consts.RED.serialize(True), 0)
        self.assertEqual(consts.BLUE.serialize(True), 1)
        self.assertEqual(consts.GREEN.serialize(True), 2)
        self.assertEqual(consts.YELLOW.serialize(True), 3)

    def testTeams(self):
        self.assertEqual(consts.Team(consts.RED).team_id, 0)
        self.assertEqual(consts.Team(consts.BLUE).team_id, 1)
        self.assertEqual(consts.Team(consts.GREEN).team_id, 2)
        self.assertEqual(consts.Team(consts.YELLOW).team_id, 3)

    def testSerializeTeams(self):
        self.assertEqual(consts.Team(consts.RED).serialize(True), 0)
        self.assertEqual(consts.Team(consts.BLUE).serialize(True), 1)
        self.assertEqual(consts.Team(consts.GREEN).serialize(True), 2)
        self.assertEqual(consts.Team(consts.YELLOW).serialize(True), 3)

if __name__ == '__main__':
    unittest.main()
