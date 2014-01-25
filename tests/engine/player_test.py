import unittest

from engine import player, consts, types, loc


class PlayerTest(unittest.TestCase):
    def testPlayerCreate(self):
        player_ = player.Player(1, consts.RED)
        self.assertEqual(player_.player_id, 1)
        self.assertEqual(player_.nation, consts.RED)

    def testPlayerSerialize(self):
        player_ = player.Player(1, consts.RED)
        self.assertEqual(player_.serialize(True), {"player_id": 1, "nation": 0, "team": 0})

if __name__ == '__main__':
    unittest.main()
