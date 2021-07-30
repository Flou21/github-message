#!/bin/bash

docker build -t registry.flou21.de/github-message .
echo "------------"
echo "------------"
echo "------------"
echo "------------"
docker run \
    -e TOKEN="****" \
    -e USERNAME="Flou21" \
    -e REPOSITORY="k8s" \
    -e MESSAGE_1="hello" \
    -e MESSAGE_2="world" \
    -e MESSAGE_3="and" \
    -e MESSAGE_4="mars" \
    -e PULL_REQUEST_NUMBER="1" \
    -e NEW_STATE="opened" \
    registry.flou21.de/github-message