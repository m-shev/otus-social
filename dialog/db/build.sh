##!/bin/bash
set -e

## build dialog db
docker build --no-cache -f ./dialog/Dockerfile -t dialog-db:latest .