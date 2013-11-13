import unittest

from engine import movement, tile


class MovementTest(unittest.TestCase):
    def testTreads(self):
        treads = movement.Treads()
        self.assertEqual(treads[tile.PLAIN], 1.0)
        self.assertEqual(treads[tile.CITY], 1.0)
        self.assertEqual(treads[tile.WOODS], 1.0)
        self.assertEqual(treads[tile.MOUNTAIN], 0.0)

    def testTires(self):
        tires = movement.Tires()
        self.assertEqual(tires[tile.PLAIN], 1.5)
        self.assertEqual(tires[tile.CITY], 1.0)
        self.assertEqual(tires[tile.WOODS], 2.0)
        self.assertEqual(tires[tile.MOUNTAIN], 0.0)

    def testFeet(self):
        feet = movement.Feet()
        self.assertEqual(feet[tile.PLAIN], 1.0)
        self.assertEqual(feet[tile.CITY], 1.0)
        self.assertEqual(feet[tile.WOODS], 1.0)
        self.assertEqual(feet[tile.MOUNTAIN], 2.0)

    def testAssignment(self):
        feet = movement.Feet()
        feet[tile.PLAIN] = 2.0
        self.assertEqual(feet[tile.PLAIN], 2.0)

    def testTireEqualsTire(self):
        self.assertEqual(movement.Tires(), movement.Tires())

    def testTireDoesNotEqualTreads(self):
        self.assertNotEqual(movement.Treads(), movement.Tires())

    def testTireDoesNotEqualModifiedTire(self):
        modified = movement.Tires()
        modified[tile.PLAIN] = modified[tile.PLAIN] + 1
        self.assertNotEqual(modified, movement.Tires())

    def testChangedMapDifferentValue(self):
        modified = movement.Tires()
        modified.movement = {tile.PLAIN: 0.0}
        self.assertNotEqual(modified, movement.Tires())

    def testChangedMapSameValue(self):
        modified = movement.Tires()
        modified.movement = {
            tile.PLAIN: 1.5,
            tile.CITY: 1.0,
            tile.WOODS: 2.0,
            tile.MOUNTAIN: 0.0,
        }
        self.assertEqual(modified, movement.Tires())

if __name__ == '__main__':
    unittest.main()
