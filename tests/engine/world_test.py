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
        self.assertEqual(world_.players[0].player_id, 13)
        self.assertEqual(world_.players[1].player_id, 26)
        self.assertEqual(world_.turn_owner, world_.players[0])

    def testWorldAddUnit(self):
        world_ = world.World([13, 26])
        tank = types.new_unit('Tank', consts.RED)
        world_.add_unit(tank, loc.Loc(1, 1))
        self.assertEqual(world_.units[(1, 1)], tank)

    def testWorldAddDoubleUnit(self):
        world_ = world.World([13, 26])
        tank1 = types.new_unit('Tank', consts.RED)
        tank2 = types.new_unit('Tank', consts.BLUE)
        world_.add_unit(tank1, loc.Loc(1, 1))
        world_.add_unit(tank2, loc.Loc(1, 1))
        self.assertEqual(world_.units[(1, 1)], tank1)

    def testWorldAddUnit(self):
        world_ = world.World([13, 26])
        tank = types.new_unit('Tank', consts.RED)
        world_.add_unit(tank, loc.Loc(100, 1))
        world_.add_unit(tank, loc.Loc(1, 100))

    def testWorldGetPlayer(self):
        world_ = world.World([13, 26])
        player1 = world_.players[0]
        player2 = world_.players[1]
        self.assertEqual(world_.get_player(13), player1)
        self.assertEqual(world_.get_player(26), player2)
        self.assertRaises(Exception, world_.get_player, 3)

    def testWorldJson(self):
        test_world = world.World([13, 26])
        test_world.add_unit(types.new_unit('Tank', consts.RED), loc.Loc(1, 1))
        expected_terrain_piece = '[%s0]' % ('0, ' * 9)
        expected_terrain = '[%s%s]' % (
            (expected_terrain_piece + ', ') * 9, expected_terrain_piece)
        #TODO Remove dump load
        player1 = '{"player": {"player_id": 13, "nation": 0, "team": 0}}'
        player2 = '{"player": {"player_id": 26, "nation": 1, "team": 1}}'
        unit = ('{ "unit": { "nation": 0, "name": "Tank", "health": 10,'
                '"distance": 7, "loc": { "x": 1, "y": 1 },'
                '"movement": { "speeds": { "0": 1.0, "1": 1.0, "2": 1.0, "3": 0.0 },'
                '"name": "Treads" } } }')
        expected_json = json.loads(
            '{ "world": {"players": [%s, %s], "terrain": %s, "units": [%s]}}' % (
                player1, player2, expected_terrain, unit))
        self.maxDiff = None
        self.assertEqual(json.loads(test_world.to_json()), expected_json)

    #TODO
    def testWorldMove(self):
        world_ = world.World([13, 26])
        world_.add_unit(types.new_unit('Tank', consts.RED), loc.Loc(1, 1))
        self.assertEqual(world_.move_unit(13, 0, [loc.Loc(1, 1)]), True)
        self.assertEqual(world_.move_unit(13, 0, [loc.Loc(1, 1), loc.Loc(2, 2)]), False)
        self.assertEqual(
            world_.move_unit(13, 0, [loc.Loc(1, 1), loc.Loc(1, 2), loc.Loc(2, 2)]),
            True)

if __name__ == '__main__':
    unittest.main()
