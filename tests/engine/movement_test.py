import unittest

from engine import types


class MovementTest(unittest.TestCase):
    def testTreads(self):
        treads = types.new_movement('Treads')
        self.assertEqual(treads[types.tiles['plain']], 1.0)
        self.assertEqual(treads[types.tiles['city']], 1.0)
        self.assertEqual(treads[types.tiles['woods']], 1.0)
        self.assertEqual(treads[types.tiles['mountain']], 0.0)

    def testTires(self):
        tires = types.new_movement('Tires')
        self.assertEqual(tires[types.tiles['plain']], 1.5)
        self.assertEqual(tires[types.tiles['city']], 1.0)
        self.assertEqual(tires[types.tiles['woods']], 2.0)
        self.assertEqual(tires[types.tiles['mountain']], 0.0)

    def testFeet(self):
        feet = types.new_movement('Feet')
        self.assertEqual(feet[types.tiles['plain']], 1.0)
        self.assertEqual(feet[types.tiles['city']], 1.0)
        self.assertEqual(feet[types.tiles['woods']], 1.0)
        self.assertEqual(feet[types.tiles['mountain']], 2.0)

    def testAssignment(self):
        feet = types.new_movement('Feet')
        feet[types.tiles['plain']] = 2.0
        self.assertEqual(feet[types.tiles['plain']], 2.0)

    def testTireEqualsTire(self):
        self.assertEqual(types.new_movement('Tires'), types.new_movement('Tires'))

    def testTireDoesNotEqualTreads(self):
        self.assertNotEqual(types.new_movement('Treads'), types.new_movement('Tires'))

    def testTireDoesNotEqualModifiedTire(self):
        modified = types.new_movement('Tires')
        modified[types.tiles['plain']] = modified[types.tiles['plain']] + 1
        self.assertNotEqual(modified, types.new_movement('Tires'))

    def testChangedMapDifferentValue(self):
        modified = types.new_movement('Tires')
        modified.dictionary = {types.tiles['plain']: 0.0}
        self.assertNotEqual(modified, types.new_movement('Tires'))

    def testChangedMapSameValue(self):
        modified = types.new_movement('Tires')
        modified.dictionary = {
            types.tiles['plain']: 1.5,
            types.tiles['city']: 1.0,
            types.tiles['woods']: 2.0,
            types.tiles['mountain']: 0.0,
        }
        self.assertEqual(modified, types.new_movement('Tires'))

    def testSerialze(self):
        self.assertEqual(
            types.new_movement('Tires').serialize(True),
            {"speeds": {0: 1.5, 1: 1.0, 2: 2.0, 3: 0.0}, "name": "Tires"}
        )

if __name__ == '__main__':
    unittest.main()
