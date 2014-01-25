import unittest

from engine import move, types, loc


#TODO Add more tests here
class MoveTest(unittest.TestCase):
    def testMoveSixTreadOverTwoPlain(self):
        valid = move.valid_move(
            6,
            types.new_movement('Treads'),
            [types.tiles['plain'], types.tiles['plain']],
            [loc.Loc(1, 1), loc.Loc(1, 2)])
        self.assertTrue(valid)

    def testMoveTwoTreadOverTenPlain(self):
        invalid = move.valid_move(
            2,
            types.new_movement('Treads'),
            [types.tiles['plain'] for i in range(10)],
            [loc.Loc(1, x) for x in range(10)])
        self.assertFalse(invalid)

    def testValidMoveInvalidTiles(self):
        invalid = move.valid_move(
            2,
            types.new_movement('Treads'),
            [types.tiles['plain'] for i in range(10)],
            [loc.Loc(1, x * 2) for x in range(10)])
        self.assertFalse(invalid)

    def testGetDistance(self):
        self.assertEqual(move.get_distance(3, 2), 1)
        self.assertEqual(move.get_distance(1, 2), 1)
        self.assertEqual(move.get_distance(1, 1), 0)
        self.assertEqual(move.get_distance(1, 4), 3)
        self.assertEqual(move.get_distance(7, 4), 3)

    def testAllTilesTouch(self):
        self.assertEqual(move.assert_all_tiles_touch([loc.Loc(1, 1), loc.Loc(2, 2)]), False)
        self.assertEqual(
            move.assert_all_tiles_touch([loc.Loc(1, 1), loc.Loc(1, 2), loc.Loc(2, 2)]),
            True)


if __name__ == '__main__':
    unittest.main()
