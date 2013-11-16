import unittest

from engine import types


class TileTest(unittest.TestCase):
    def testTile(self):
        self.assertEqual(types.tiles['plain'], 0)
        self.assertEqual(types.tiles['city'], 1)
        self.assertEqual(types.tiles['woods'], 2)
        self.assertEqual(types.tiles['mountain'], 3)

if __name__ == '__main__':
    unittest.main()
