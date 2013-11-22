#!/usr/bin/env python

import os
import json
import argparse

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


def delete_all_templates():
    do_delete = raw_input('Print remove contents of %s? (y/n) ' % (template_dir,))
    if do_delete == 'y':
        multi_map(delete_template, structs)
        os.rmdir(template_dir)
    else:
        print 'Aborting on user request'


def delete_template(new_, name, args):
    os.remove(os.path.join(template_dir, '%s.json' % (name,)))


def generate_template(new_, name, args):
    with open(os.path.join(template_dir, '%s.json' % (name,)), 'w') as f:
        f.write(
            without_trailing_whitespace(
                json.dumps(
                    json.loads(
                        repr(
                            new_(name, *args))),
                    indent=4)))


def generate_all_templates():
    if not os.path.exists(template_dir):
        os.mkdir(template_dir)
    multi_map(generate_template, structs)


def main():
    parser = argparse.ArgumentParser()
    subparsers = parser.add_subparsers(help='sub-command help')

    delete_parser = subparsers.add_parser(
        'delete', help='delete all template files')
    delete_parser.set_defaults(func=delete_all_templates)

    generate_parser = subparsers.add_parser(
        'generate', help='generate all template files')
    generate_parser.set_defaults(func=generate_all_templates)

    args = parser.parse_args()
    args.func()

if __name__ == '__main__':
    main()
