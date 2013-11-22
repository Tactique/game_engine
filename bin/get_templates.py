#!/usr/bin/env python

import json
import os
from lib.functional import multi_map

from engine import types, consts

template_dir = os.path.join(os.environ['PORTER'], 'templates')

structs = (
    (types.new_unit, "Tank", (consts.RED,)),
    (types.new_attack, "RegularCannon", ()),
    (types.new_armor, "WeakMetal", ()),
    (types.new_movement, "Treads", ()),
)


def without_trailing_whitespace(string):
    def remove_trailing_whitespace(line):
        return line.rstrip()

    return '\n'.join(map(remove_trailing_whitespace, string.split('\n')))


def generate_template(new_, name, args):
    with open(os.path.join(template_dir, '%s.json' % (name,)), 'w') as f:
        f.write(
            without_trailing_whitespace(
                json.dumps(
                    json.loads(
                        repr(
                            new_(name, *args))),
                    indent=4)))


def main():
    if not os.path.exists(template_dir):
        os.mkdir(template_dir)
    multi_map(generate_template, structs)

if __name__ == '__main__':
    main()
