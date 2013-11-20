import unittest

from engine import file_loader


class FileLoaderTest(unittest.TestCase):
    def test_load_units(self):
        dicts = file_loader.read_and_parse_json('units')
        self.assertIsInstance(dicts, list)
        self.assertGreater(len(dicts), 0)
        for dict_ in dicts:
            self.assertIsInstance(dict_, dict)

    def testLoadStruct(self):
        unit_map = file_loader.load_struct('units')
        for unit_name, unit_args in unit_map.items():
            self.assertIsInstance(unit_name, str)
            self.assertIsInstance(unit_args, dict)

    def testLoadEnum(self):
        unit_map = file_loader.load_enum('attack_types')
        self.assertIsInstance(unit_map, dict)
        for unit_name, unit_enum in unit_map.items():
            self.assertIsInstance(unit_name, str)
            self.assertIsInstance(unit_enum, int)

if __name__ == '__main__':
    unittest.main()
