import unittest
#TODO Remove jsoning here
import json

from engine import world, player, consts, types, loc


#TODO Add more tests here
class WorldTest(unittest.TestCase):
    def testWorldCreate(self):
        world_ = world.World([13, 26])
        self.assertEqual(
            world_.terrain,
            [[types.tiles['plain'] for i in range(10)] for i in range(10)])
        self.assertEqual(world_.players[consts.RED].units, {})
        self.assertEqual(world_.players[consts.RED].player_id, 13)
        self.assertEqual(world_.players[consts.BLUE].units, {})
        self.assertEqual(world_.players[consts.BLUE].player_id, 26)
        self.assertEqual(world_.turn_owner, world_.players[consts.RED])

    def testWorldAddUnit(self):
        world_ = world.World([13, 26])
        tank = types.new_unit('Tank', consts.RED, loc.Loc(1, 1))
        world_.add_unit(26, tank)
        self.assertEqual(world_.get_player(26).units[0], tank)

    def testWorldGetPlayer(self):
        world_ = world.World([13, 26])
        player1 = world_.players[consts.RED]
        player2 = world_.players[consts.BLUE]
        self.assertEqual(world_.get_player(13), player1)
        self.assertEqual(world_.get_player(26), player2)
        self.assertRaises(Exception, world_.get_player, 3)

    def testWorldJson(self):
        test_world = world.World([13, 26])
        expected_terrain_piece = '[%s0]' % ('0, ' * 9)
        expected_terrain = '[%s%s]' % (
            (expected_terrain_piece + ', ') * 9, expected_terrain_piece)
        #TODO Remove dump load
        expected_json = json.loads(
            '{"13": {}, "26": {}, "terrain": %s}' % expected_terrain)
        self.assertEqual(json.loads(test_world.to_json()), expected_json)

    #TODO
    def testWorldMove(self):
        world_ = world.World([13, 26])
        world_.add_unit(13, types.new_unit('Tank', consts.RED, loc.Loc(1, 1)))
        self.assertEqual(world_.move_unit(13, 0, [(1, 1)]), True)
        self.assertEqual(world_.move_unit(13, 0, [(1, 1), (2, 2)]), False)
        self.assertEqual(
            world_.move_unit(13, 0, [(1, 1), (1, 2), (2, 2)]),
            True)

if __name__ == '__main__':
    unittest.main()
