import unittest

from lib import contract


class AcceptsTest(unittest.TestCase):
    def testAcceptsInt(self):
        @contract.accepts(int)
        def int_taker(number):
            return number
        self.assertEqual(int_taker(5), 5)
        self.assertRaises(contract.ContractBrokenError, int_taker, 'boo')
        self.assertRaises(contract.ContractBrokenError, int_taker, 'boo')

    def testAcceptsString(self):
        @contract.accepts(str)
        def string_taker(string):
            return string
        self.assertEqual(string_taker('boo'), 'boo')
        self.assertRaises(contract.ContractBrokenError, string_taker, 5)

    def testAcceptsStringAndInt(self):
        @contract.accepts(str, int)
        def string_int_taker(string, number):
            return '%s %s' % (string, number)
        self.assertEqual(string_int_taker('boo', 5), 'boo 5')
        self.assertRaises(contract.ContractBrokenError, string_int_taker, 5)

    def testAcceptsNothing(self):
        @contract.accepts()
        def nothing_taker():
            return 1
        self.assertEqual(nothing_taker(), 1)

    def testAcceptsFakeNothing(self):
        @contract.accepts()
        def fake_nothing_taker(fake_number):
            pass
        self.assertRaises(contract.ContractBrokenError, fake_nothing_taker, 5)

    def testAcceptsCustomNewStyleClass(self):
        class NewStyleClass(object):
            1

        class OtherNewStyleClass(object):
            2

        @contract.accepts(NewStyleClass)
        def new_taker(new_class):
            return new_class
        new_class = NewStyleClass()
        new_other = OtherNewStyleClass()
        self.assertEqual(new_taker(new_class), new_class)
        self.assertRaises(contract.ContractBrokenError, new_taker, new_other)

    def testAcceptsCustomSuperClass(self):
        class NewStyleClass(object):
            1

        class SubNewStyleClass(NewStyleClass):
            2

        @contract.accepts(NewStyleClass)
        def new_taker(new_class):
            return new_class
        my_new_sub = SubNewStyleClass()
        my_new_class = NewStyleClass()
        self.assertEqual(new_taker(my_new_sub), my_new_sub)
        self.assertEqual(new_taker(my_new_class), my_new_class)
        self.assertRaises(contract.ContractBrokenError, new_taker, 1)

    def testSelfAccepts(self):
        class NewStyleClass(object):
            @contract.self_accepts(int)
            def return_input_int(self, number):
                return number

        new_class = NewStyleClass()
        self.assertEqual(new_class.return_input_int(1), 1)

    def testAcceptsList(self):
        @contract.accepts([int])
        def accepts_int_list(ints):
            return ints

        self.assertEqual(accepts_int_list([1, 2, 3]), [1, 2, 3])

    def testAcceptsTuple(self):
        @contract.accepts((int,))
        def accepts_int_tuple(ints):
            return ints

        self.assertEqual(accepts_int_tuple((1, 2, 3)), (1, 2, 3))

    def testAcceptsDict(self):
        @contract.accepts({str: int})
        def accepts_str_to_int_dict(str_to_int):
            return str_to_int

        self.assertEqual(accepts_str_to_int_dict({'hi': 0, 'bye': 1}), {'hi': 0, 'bye': 1})


class ReturnTest(unittest.TestCase):
    def testReturnsAsPromised(self):
        @contract.returns(int)
        def returns_int():
            return 1
        ret = returns_int()
        self.assertEqual(ret, 1)

    def testReturnsRaisesIncorrectType(self):
        @contract.returns(int)
        def doesnt_return_int():
            return 'hi'
        self.assertRaises(contract.ContractBrokenError, doesnt_return_int)

    def testReturnsNothing(self):
        @contract.returns(None)
        def returns_nothing():
            pass
        ret = returns_nothing()
        self.assertEqual(ret, None)

    def testDoesntActuallyReturnNothing(self):
        @contract.returns(None)
        def doesnt_return_nothing():
            return 'hi'
        self.assertRaises(contract.ContractBrokenError, doesnt_return_nothing)


if __name__ == '__main__':
    unittest.main()
