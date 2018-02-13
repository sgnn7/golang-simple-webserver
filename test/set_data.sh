#!/bin/bash -e

CURRENT_DIR=$(dirname $0)
SERVER=localhost:8000

curl -X POST http://$SERVER/api/objects -d @$CURRENT_DIR/data.json \
          --header "Content-Type: application/json"
