#!/usr/bin/env bash

REPODIR=$(cd "$(dirname $0)/.."; pwd)
cd $REPODIR
source script/path_info.cfg
source script/log.sh
mkdir -p run/
go mod tidy

start_service() {
    service="$1"
    service_name="$2"
    service_dir="$3"
    cd $service_dir
    make build
    ./${service} > $REPODIR/run/${service}.log 2>&1 &
    pid=$!
    echo "$pid" > $REPODIR/run/${service}.pid
    log_info "${service} [$pid] started"
}

service="$1"

# start all
if [ -z "$service" ]; then
    for ((i = 0; i < ${#services[*]}; i++)); do
        service=${services[$i]}
        service_name=${service_names[$i]}
        service_dir=${service_dirs[$i]}
        start_service $service $service_name $service_dir
    done
    exit 0
fi

# start single
for ((i = 0; i < ${#services[*]}; i++)); do
    s=${services[$i]}
    if [ "$s" = "$service" ]; then
        service_name=${service_names[$i]}
        service_dir=${service_dirs[$i]}
    fi
done
if [ -z "$service_names" ] || [ -z "$service_dir" ]; then
    log_error "service not found: $service"
    exit -1
fi

start_service $service $service_name $service_dir
