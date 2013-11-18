import os
import json

data_dir = os.path.join(os.environ['PORTER'], 'data')


def read_and_parse_json(data_type):
    sub_dir = os.path.join(data_dir, data_type)

    def full_path(file_name):
        return os.path.join(sub_dir, file_name)

    def only_json(file_name):
        return file_name.endswith('.json')

    def load_json(json_file_name):
        with open(json_file_name) as json_file:
            return json.load(json_file)

    return map(load_json, filter(only_json, map(full_path, os.listdir(sub_dir))))
