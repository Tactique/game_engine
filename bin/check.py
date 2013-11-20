#!/usr/bin/env python

# Assumes env variable 'PORTER' is set to where the 'porter' dir exists

import os

import pep8

from util import find_all
from verify_data import verify_all
from coverage_check import coverage_test_all


def pep8_all():
    def pep8_file(file_):
        checker = pep8.Checker(filename=file_, max_line_length=99)
        incorrect = checker.check_all()
        if incorrect != 0:
            raise Exception("pep8 on file %s failed" % (file_,))
        print '%s pep8 compliant' % (file_,)

    map(pep8_file, find_all(os.environ['PORTER'], '.py'))


def main():
    os.chdir(os.environ['PORTER'])
    pep8_all()
    verify_all()
    coverage_test_all()
    print 'OK'

if __name__ == '__main__':
    main()
