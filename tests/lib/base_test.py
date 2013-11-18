import unittest

from lib import base


class BaseTest(unittest.TestCase):
    def testBaseClass(self):
        class SuperBase(base.BaseClass):
            def __init__(self, a):
                self.a = a

        class OtherSuperBase(base.BaseClass):
            def __init__(self, a):
                self.a = a

        s1 = SuperBase(1)
        s2 = SuperBase(2)
        s3 = SuperBase(1)
        os1 = OtherSuperBase(1)
        self.assertEqual(s1, s1)
        self.assertEqual(s1, s3)
        self.assertNotEqual(s1, s2)
        self.assertNotEqual(s1, os1)

    def testBaseEquality(self):
        class SuperBig(base.BaseClass):
            def __init__(self, a, b, c):
                self.a = a
                self.b = b
                self.c = c

        class OtherSuperBig(base.BaseClass):
            def __init__(self, a, b, c):
                self.a = a
                self.b = b
                self.c = c

        s1 = SuperBig(1, 2, 3)
        s2 = SuperBig(1, 2, 3)
        s3 = SuperBig(2, 2, 3)
        os1 = OtherSuperBig(1, 2, 3)

        self.assertEqual(s1, s2)
        self.assertNotEqual(s1, s3)
        self.assertNotEqual(os1, s1)
        self.assertNotEqual(os1, s1)
        self.assertNotEqual(os1, s3)

    def testBaseDictionary(self):
        class Diction(base.BaseDictionary):
            def __init__(self):
                base.BaseDictionary.__init__(self, {0: 1.0, 1: 2.0})

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

    def testToStriong(self):
        class ThreeAttrs(base.BaseClass):
            def __init__(self):
                self.a = 1
                self.b = 2
                self.c = 3

        self.assertEqual(ThreeAttrs().__repr__(), 'a:1 b:2 c:3 ')

if __name__ == '__main__':
    unittest.main()
