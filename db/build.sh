##!/bin/bash
set -e
source ".msql"

## build master
#docker build --no-cache --build-arg sec="$R" -t sigma-social-db:latest .

## build slave1
docker build --no-cache --build-arg sec="$R" -f ./slave1/Dockerfile -t sigma-social-db-replica-1:latest .