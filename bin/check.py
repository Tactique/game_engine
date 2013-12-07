#!/usr/bin/env python

# Assumes env variable 'PORTER' is set to where the 'porter' dir exists
# Assumes PYTHONPATH is set to 'src' dir under 'PORTER'

import os

import color

import data_integrity_check
import coverage_check
import pep8_check


def main():
    pep8_check.pep8_all()
    data_integrity_check.verify_all()
    coverage_check.coverage_test_all()
    print color.colorize('OK', color.green)

if __name__ == '__main__':
    main()
