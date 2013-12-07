#!/usr/bin/env python

import os
import argparse


def main():
    project_path = os.path.split(
        os.path.dirname(
            os.path.realpath(__file__)))[0]
    export_project_path = 'export PORTER=%s' % (
        project_path,)
    export_python_path = 'export PYTHONPATH=$PYTHONPATH:$PORTER/%s' % (
        'src',)

    parser = argparse.ArgumentParser()
    parser.add_argument(
        '-e', '--export-only',
        action='store_true',
        help='Only print the export lines')

    args = parser.parse_args()
    if args.export_only:
        print export_project_path
        print export_python_path
    else:
        print 'Setting up...'
        print 'Either run the following or add it to your .bashrc or equivalent:'
        print export_project_path
        print export_python_path


if __name__ == '__main__':
    main()
