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

if __name__ == '__main__':
    unittest.main()
