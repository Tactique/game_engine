import os
import json

from lib import contract

data_dir = os.path.join(os.environ['PORTER'], 'data')


@contract.accepts(str)
@contract.returns(list)
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


@contract.accepts(str)
@contract.returns(dict)
def load_enum(struct_name):
    enum_ = {}
    list_ = read_and_parse_json(struct_name)[0]
    for enumeration, enum_type in enumerate(list_):
        enum_[str(enum_type)] = enumeration
    return enum_


@contract.accepts(str)
@contract.returns(dict)
def load_struct(struct_name):
    def create_struct_map(struct_map, struct_):
        struct_map[str(struct_['name'])] = struct_
        return struct_map

    return reduce(create_struct_map, read_and_parse_json(struct_name), {})
