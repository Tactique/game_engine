import os
import json

data_dir = os.path.join(os.environ['PORTER'], 'data')


def read_and_parse_json(data_type):
    sub_dir = os.path.join(data_dir, data_type)
    elements = []

    def full_path(file_name):
        return os.path.join(sub_dir, file_name)

    def only_json(file_name):
        return file_name.endswith('.json')

    for json_file_name in filter(only_json, map(full_path, os.listdir(sub_dir))):
        with open(json_file_name) as json_file:
            elements.append(json.load(json_file))
    return elements
