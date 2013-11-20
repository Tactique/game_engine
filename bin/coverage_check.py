#!/usr/bin/env python

import os
import subprocess

from lib import functional

from util import find_all


def coverage_module(package, module):
    command = (
        'coverage run --branch'
        ' --source=%s.%s tests/%s/%s_test.py')
    print subprocess.check_output(
        command % (package, module, package, module),
        stderr=subprocess.STDOUT,
        shell=True)
    print subprocess.check_output(
        'coverage report --fail-under=100 -m',
        stderr=subprocess.STDOUT,
        shell=True)
    subprocess.check_output(
        'coverage erase',
        shell=True)


def coverage_test_package(package):
    def path_to_name(name):
        return os.path.split(name)[1].split('.')[0]

    for module in functional.removed(
            map(path_to_name, find_all(
                os.path.join('src', package), '.py')), '__init__'):
        print package, module
        coverage_module(package, module)


def coverage_test_all():
    for package in ['lib', 'engine']:
        coverage_test_package(package)
