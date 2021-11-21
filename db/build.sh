##!/bin/bash
set -e
source "common/.msql"

commonConfig="
# Путь к relay логу
relay-log = /var/log/mysql/mysql-relay-bin.log

# Путь к bin логу на Мастере
log_bin = /var/log/mysql/mysql-bin.log

# название Вашей базы данных, которая будет реплицироваться
binlog_do_db=socialdb

### Error log settings
log_error	= /var/log/mysql/error.log

### General log settings
general_log = 1
general_log_file = /var/log/mysql/general.log"

mycnfReplica1="[mysqld]

# ID slave
server-id = 2
"

mycnfReplica2="[mysqld]
# ID slave
server-id = 3"

## build master
docker build --no-cache --build-arg sec="$R" -t sigma-social-db:latest .

# build slave1
rm -f ./slave/.my.cnf

echo "$mycnfReplica1 $commonConfig" >> ./slave/.my.cnf
docker build --no-cache --build-arg sec="$R" -f ./slave/Dockerfile -t sigma-social-db-replica-1:latest .

## build slave2
rm -f ./slave/.my.cnf
echo "$mycnfReplica1 $commonConfig" >> ./slave/.my.cnf
docker build --no-cache --build-arg sec="$R" -f ./slave/Dockerfile -t sigma-social-db-replica-2:latest .

#build proxysql
docker build --no-cache -f ./proxysql/Dockerfile -t sigma-db-proxy:latest .