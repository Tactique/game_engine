import unittest

from engine import move, movement, tile


#TODO Add more tests here
class MoveTest(unittest.TestCase):
    def testMoveSixTreadOverTwoPlain(self):
        valid = move.valid_move(
            6, movement.Treads(), [tile.PLAIN, tile.PLAIN])
        self.assertTrue(valid)

    def testMoveTwoTreadOverTenPlain(self):
        invalid = move.valid_move(
            2, movement.Treads(), [tile.PLAIN for i in range(10)])
        self.assertFalse(invalid)

if __name__ == '__main__':
    unittest.main()
