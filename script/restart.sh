#!/usr/bin/env bash

REPODIR=$(cd "$(dirname $0)/.."; pwd)
cd $REPODIR

bash ./script/stop.sh $@

rm ./run/*.log
rm ./run/*.pid

bash ./script/start.sh $@
