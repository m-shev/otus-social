##!/bin/bash
set -e
source ".msql"
docker build --no-cache --build-arg sec="$R" -t sigma-social-db:latest .