#!/usr/bin/env bash
SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)
source ${SCRIPT_DIR}/bash_functions.sh

http -v GET http://localhost:8080/event/$(sed 's|\/|_|g' ${SCRIPT_DIR}/sqs_request_handle.dat)
