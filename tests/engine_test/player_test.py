import unittest

from engine import player, unit


class PlayerTest(unittest.TestCase):
    def testPlayerCreate(self):
        player_ = player.Player(1)
        self.assertEqual(player_.uid, 1)
        self.assertEqual(player_.units, [])

    def testPlayerAddUnit(self):
        player_ = player.Player(1)
        tank = unit.Tank()
        player_.add_unit(tank)
        self.assertEqual(player_.units, [tank])

if __name__ == '__main__':
    unittest.main()
