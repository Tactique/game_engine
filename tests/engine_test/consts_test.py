import unittest

from engine import consts


class ConstsTest(unittest.TestCase):
    def test_max_health(self):
        self.assertEqual(consts.MAX_HEALTH, 10)

if __name__ == '__main__':
    unittest.main()
