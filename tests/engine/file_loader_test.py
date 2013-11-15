import unittest

from engine import file_loader


class FileLoaderTest(unittest.TestCase):
    def test_load_units(self):
        dicts = file_loader.read_and_parse_json('units')
        self.assertIsInstance(dicts, list)
        self.assertGreater(len(dicts), 0)
        for dict_ in dicts:
            self.assertIsInstance(dict_, dict)

if __name__ == '__main__':
    unittest.main()
