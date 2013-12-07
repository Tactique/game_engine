#!/usr/bin/env python

import os

from lib import functional

from util import find_all


def coverage_module(package, module):
    command = (
        'coverage run --branch'
        ' --source=%s.%s tests/%s/%s_test.py') % (
        package, module, package, module)
    out, ret = util.check_call_output(command, stderr=util.STDOUT, shell=True)
    print out
    out, ret = util.check_call_output(
        'coverage report --fail-under=100 -m', stderr=util.STDOUT, shell=True)
    print out
    util.check_call_output('coverage erase', shell=True)


def coverage_test_package(package):
    def path_to_name(name):
        return os.path.split(name)[1].split('.')[0]

    for module in functional.removed(
            map(path_to_name, find_all(
                os.path.join('src', package), '.py')), '__init__'):
        print package, module
        coverage_module(package, module)


def coverage_test_all():
    os.chdir(os.environ['PORTER'])
    for package in os.listdir('src/'):
        coverage_test_package(package)

if __name__ == '__main__':
    coverage_test_all()
