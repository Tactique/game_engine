import unittest

from engine import move, movement, tile


#TODO Add more tests here
class MoveTest(unittest.TestCase):
    def testMoveSixTreadOverTwoPlain(self):
        valid = move.valid_move(
            6, movement.Treads(), [tile.PLAIN, tile.PLAIN], [(1, 1), (1, 2)])
        self.assertTrue(valid)

    def testMoveTwoTreadOverTenPlain(self):
        invalid = move.valid_move(
            2,
            movement.Treads(),
            [tile.PLAIN for i in range(10)],
            [(1, x) for x in range(10)])
        self.assertFalse(invalid)

    def testValidMoveInvalidTiles(self):
        invalid = move.valid_move(
            2,
            movement.Treads(),
            [tile.PLAIN for i in range(10)],
            [(1, x * 2) for x in range(10)])
        self.assertFalse(invalid)

    def testGetDistance(self):
        self.assertEqual(move.get_distance(3, 2), 1)
        self.assertEqual(move.get_distance(1, 2), 1)
        self.assertEqual(move.get_distance(1, 1), 0)
        self.assertEqual(move.get_distance(1, 4), 3)
        self.assertEqual(move.get_distance(7, 4), 3)

    def testAllTilesTouch(self):
        self.assertEqual(move.assert_all_tiles_touch([(1, 1), (2, 2)]), False)
        self.assertEqual(
            move.assert_all_tiles_touch([(1, 1), (1, 2), (2, 2)]),
            True)


if __name__ == '__main__':
    unittest.main()
