#!/bin/sh

pushd $(dirname $0) > /dev/null

kill $(pgrep -f "python -u main.py")
bin/setup.py -e > .env.tmp
source ./.env.tmp 
rm .env.tmp
python -u main.py > server.log 2>&1 &

popd > /dev/null
