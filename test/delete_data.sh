#!/bin/bash -e

CURRENT_DIR=$(dirname $0)
SERVER=localhost:8000

curl -X DELETE http://$SERVER/api/objects/my-document
echo "OK"
