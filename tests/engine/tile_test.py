import unittest

from engine import tile


class TileTest(unittest.TestCase):
    def testTile(self):
        self.assertEqual(tile.PLAIN, 0)
        self.assertEqual(tile.CITY, 1)
        self.assertEqual(tile.WOODS, 2)
        self.assertEqual(tile.MOUNTAIN, 3)

if __name__ == '__main__':
    unittest.main()
