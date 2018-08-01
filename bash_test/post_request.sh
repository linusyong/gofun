#!/usr/bin/env bash
SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)

http POST http://localhost:8080/event < ${SCRIPT_DIR}/request.json | jq -r '.status' | jq -r '.webseal[]'
