import unittest

from engine import loc


class LocTest(unittest.TestCase):
    def testLocs(self):
        self.assertEqual(loc.Loc(1, 0).x, 1)
        self.assertEqual(loc.Loc(1, 0).y, 0)
        self.assertEqual(loc.Loc(2, 5).x, 2)
        self.assertEqual(loc.Loc(5, 2).y, 2)

    def testLocSerialize(self):
        self.assertEqual(
            loc.Loc(1, 0).serialize(True),
            {'x': 1, 'y': 0}
        )

if __name__ == '__main__':
    unittest.main()
