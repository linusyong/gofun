#!/usr/bin/env bash
curl -sX POST http://localhost:8080/event -d @request.json| jq -r '.status' | jq -r '.webseal[]'
