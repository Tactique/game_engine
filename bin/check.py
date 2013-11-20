#!/usr/bin/env python

# Assumes env variable 'PORTER' is set to where the 'porter' dir exists

import os

from verify_data import verify_all
from coverage_check import coverage_test_all
from pep8_check import pep8_all


def main():
    pep8_all()
    verify_all()
    coverage_test_all()
    print 'OK'

if __name__ == '__main__':
    main()
