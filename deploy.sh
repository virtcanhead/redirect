#!/bin/bash

set -e
set -u

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker build -t canhead/redirect .
docker push canhead/redirect
