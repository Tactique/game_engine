#!/usr/bin/env python

import os
import sys
import subprocess
import unittest
import coverage

import pep8

from lib import functional


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


def coverage_module(package, module):
    '''
    cover = coverage.coverage(branch=True, timid=True, source=['%s/%s.py' % (package, module,)])
    cover.start()

    suite = unittest.defaultTestLoader.loadTestsFromName(
        'tests.%s_test.%s_test' % (package, module,))

    unittest.runner.TextTestRunner(descriptions=True, verbosity=1).run(suite)
    cover.stop()

    percent = cover.report()
    print percent
    assert(percent == 100.0)

    '''
    command = (
        'coverage run --branch'
        ' --source=%s.%s tests/%s_test/%s_test.py')
    print subprocess.check_output(
        command % (package, module, package, module),
        stderr=subprocess.STDOUT,
        shell=True)
    print subprocess.check_output(
        'coverage report --fail-under=100 -m',
        stderr=subprocess.STDOUT,
        shell=True)


def coverage_test_package(package):
    def path_to_name(name):
        return name.split('/')[-1].split('.')[0]

    for module in functional.removed(map(path_to_name, find_all(package, '.py')), '__init__'):
        print package, module
        coverage_module(package, module)


def coverage_test_all():
    for package in ['lib', 'engine']:
        coverage_test_package(package)


def main():
    pep8_all()
    coverage_test_all()
    print 'OK'

if __name__ == '__main__':
    main()
