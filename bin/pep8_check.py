#!/usr/bin/env python

import os
import pep8

from util import find_all


def pep8_all():
    def pep8_file(file_):
        checker = pep8.Checker(filename=file_, max_line_length=99)
        incorrect = checker.check_all()
        if incorrect != 0:
            raise Exception("pep8 on file %s failed" % (file_,))
        print '%s pep8 compliant' % (file_,)

    map(pep8_file, find_all(os.environ['PORTER'], '.py'))

if __name__ == '__main__':
    pep8_all()
