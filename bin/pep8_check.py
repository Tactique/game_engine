#!/usr/bin/env python

import os

import pep8

import notify
import util

notifier = notify.Notifier('Pep8')


def pep8_all():
    def pep8_file(file_):
        checker = pep8.Checker(filename=file_, max_line_length=99)
        incorrect = checker.check_all()
        if incorrect != 0:
            notifier.failure(str(file_))
        notifier.success(str(file_))

    map(pep8_file, util.find_all(os.environ['PORTER'], '.py'))

if __name__ == '__main__':
    pep8_all()
