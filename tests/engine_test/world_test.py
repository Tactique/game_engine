import unittest

from engine import world, player, unit, tile


#TODO Add more tests here
class WorldTest(unittest.TestCase):
    def testWorldCreate(self):
        world_ = world.World([1, 2])
        self.assertEqual(
            world_.terrain,
            [[tile.PLAIN for i in range(10)] for i in range(10)])
        self.assertEqual(world_.players['Red'].units, [])
        self.assertEqual(world_.players['Red'].uid, 1)
        self.assertEqual(world_.players['Blue'].units, [])
        self.assertEqual(world_.players['Blue'].uid, 2)

    def testWorldAddUnit(self):
        world_ = world.World([1, 2])
        tank = unit.Tank()
        world_.add_unit(1, tank)
        self.assertEqual(world_.get_player(1).units[0], tank)

    def testWorldGetPlayer(self):
        world_ = world.World([1, 2])
        player1 = world_.players['Red']
        player2 = world_.players['Blue']
        self.assertEqual(world_.get_player(1), player1)
        self.assertEqual(world_.get_player(2), player2)
        self.assertRaises(Exception, world_.get_player, 3)

    def testWorldJson(self):
        test_world = world.World([1, 2])
        expected_json_piece = '[%s0]' % ('0, ' * 9)
        expected_json = '[%s%s]' % (
            (expected_json_piece + ', ') * 9, expected_json_piece)
        self.assertEqual(test_world.to_json(), expected_json)

    #TODO
    def testWorldMove(self):
        test_world = world.World([1, 2])
        self.assertEqual(test_world.move(1, (1, 2), []), False)

if __name__ == '__main__':
    unittest.main()
