import unittest

from lib import functional


class FunctionalTest(unittest.TestCase):
    def testRemoved(self):
        a = [1, 2, 3, 4]
        b = functional.removed(a, 3)
        self.assertEqual(a, [1, 2, 3, 4])
        self.assertEqual(b, [1, 2, 4])

    def testRemovedMultiple(self):
        a = [1, 2, 3, 3, 4]
        b = functional.removed(a, 3)
        self.assertEqual(a, [1, 2, 3, 3, 4])
        self.assertEqual(b, [1, 2, 4])

    def testRemovedNone(self):
        a = [1, 2, 3, 4]
        b = functional.removed(a, 5)
        self.assertEqual(a, [1, 2, 3, 4])
        self.assertEqual(b, [1, 2, 3, 4])

    def testRemovedMismatchedType(self):
        a = [1, 2, 3, 4]
        b = functional.removed(a, 'a')
        self.assertEqual(a, [1, 2, 3, 4])
        self.assertEqual(b, [1, 2, 3, 4])

    def testMultiMap(self):
        def accepts_two(x, y):
            return x + y

        self.assertEqual(functional.multi_map(accepts_two, enumerate([10, 20, 30])), [10, 21, 32])

    def testMultiReduce(self):
        def adds_x_and_y(val, x, y):
            return val + x + y

        self.assertEqual(functional.multi_reduce(adds_x_and_y, [(1, 2), (3, 4)], 0), 10)

if __name__ == '__main__':
    unittest.main()
