#!/bin/bash -e

CURRENT_DIR=$(dirname $0)
SERVER=localhost:8000

curl -s http://$SERVER/api/objects/my-document | jq
echo "OK"
