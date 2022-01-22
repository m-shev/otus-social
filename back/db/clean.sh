##!/bin/bash
set -e

docker volume rm social-db-replica-1-log
docker volume rm social-db-replica-1-data

docker volume rm social-db-replica-2-log
docker volume rm social-db-replica-2-data

#docker volume rm social-db-data
#docker volume rm social-db-log
