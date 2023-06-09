#!/bin/bash

#Shell font formatting information
COLOR_SUFFIX="\033[0m"
BLACK_PREFIX="\033[30m"
RED_PREFIX="\033[31m"
GREEN_PREFIX="\033[32m"
YELLOW_PREFIX="\033[33m"
BLUE_PREFIX="\033[34m"
PURPLE_PREFIX="\033[35m"
SKY_BLUE_PREFIX="\033[36m"


log_error() {
    echo -e "${RED_PREFIX}$@${COLOR_SUFFIX}"
}

log_info() {
    echo -e "${GREEN_PREFIX}$@${COLOR_SUFFIX}"
}

log_debug() {
    echo -e "${BLACK_PREFIX}$@${COLOR_SUFFIX}"
}


