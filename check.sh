#!/bin/bash

#set -e

function check_ok {
    if [ $? -ne 0 ]; then
        echo $1
        echo "FAILED"
        exit 1
    fi
}

find . -name "*.py" | xargs pep8
check_ok "pep8"

for package in 'lib' 'engine'; do
    for module in ${package}/*; do
        module=$(echo ${module} | sed "s/${package}\/\(.*\).py/\1/")
        echo ${module}
        if [ ${module} != '__init__' ]; then
            coverage run --branch --source=${package}.${module} tests/${package}_test/${module}_test.py
            check_ok "Unit Test"
            coverage report --fail-under=100 -m
            check_ok "Code Coverage"
            coverage erase
        fi
    done
done
echo "All Done:"
echo "OK"
