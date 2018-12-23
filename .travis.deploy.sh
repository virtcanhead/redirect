#!/bin/bash

set -e
set -u

echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o redirect
docker build -t canhead/redirect .
docker push canhead/redirect
