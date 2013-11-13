import unittest

from engine import base


class BaseTest(unittest.TestCase):
    def testBaseClass(self):
        base_1 = base.BaseClass()
        base_1.a = 1
        base_2 = base.BaseClass()
        base_2.a = 2
        base_3 = base.BaseClass()
        base_3.a = 1
        self.assertNotEqual(base_1, base_2)
        self.assertNotEqual(base_1, base_3)
        self.assertEqual(base_1, base_1)

    def testBaseDictionary(self):
        class Diction(base.BaseDictionary):
            def __init__(self):
                self.dictionary = {0: 1.0, 1: 2.0}

        regular = Diction()
        regular_2 = Diction()
        all_double = Diction()
        all_double[0] = 2.0

        self.assertEqual(regular, regular_2)
        self.assertNotEqual(regular, all_double)
        self.assertEqual(regular[0], 1.0)
        self.assertEqual(regular[1], 2.0)
        self.assertEqual(all_double[0], 2.0)
        self.assertEqual(all_double[1], 2.0)

if __name__ == '__main__':
    unittest.main()
