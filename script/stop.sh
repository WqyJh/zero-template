#!/usr/bin/env bash

REPODIR=$(cd "$(dirname $0)/.."; pwd)
cd $REPODIR
source script/path_info.cfg
source script/log.sh

stop_service() {
    service="$1"
    service_name="$2"
    service_dir="$3"
    pid=$(cat $REPODIR/run/${service}.pid)
    kill $pid
    log_info "${service} [$pid] stopped"
}

service="$1"

# stop all
if [ -z "$service" ]; then
    for ((i = 0; i < ${#services[*]}; i++)); do
        service=${services[$i]}
        service_name=${service_names[$i]}
        service_dir=${service_dirs[$i]}
        stop_service $service $service_name $service_dir
    done
    exit 0
fi

# stop single
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

stop_service $service $service_name $service_dir
