#!/usr/bin/env python

# Assumes env variable 'PORTER' is set to where the 'porter' dir exists

import os
import subprocess

import pep8

from lib import functional
from verify_data import verify_all


def find_all(directory, pattern):
    files = []

    def prepend_directory(file_):
        return os.path.join(directory, file_)
    for file_ in map(prepend_directory, os.listdir(directory)):
        if os.path.isdir(file_):
            files.extend(find_all(file_, pattern))
        elif os.path.isfile(file_) and file_.endswith(pattern):
            files.append(file_)
    return files


def pep8_all():
    for file_ in find_all('.', '.py'):
        checker = pep8.Checker(filename=file_, max_line_length=99)
        incorrect = checker.check_all()
        if incorrect != 0:
            raise Exception("pep8 on file %s failed" % (file_,))
        print '%s pep8 compliant' % (file_,)


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


def main():
    verify_all()
    os.chdir(os.environ['PORTER'])
    pep8_all()
    coverage_test_all()
    print 'OK'

if __name__ == '__main__':
    main()
