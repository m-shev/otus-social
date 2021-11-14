##!/bin/bash
set -e
source ".msql"

mycnf1="[mysqld]
# ID slave
server-id = 2

# Путь к relay логу
relay-log = /var/log/mysql/mysql-relay-bin.log

# Путь к bin логу на Мастере
log_bin = /var/log/mysql/mysql-bin.log

# название Вашей базы данных, которая будет реплицироваться
binlog_do_db=socialdb"

mycnf2="[mysqld]
       # ID slave
       server-id = 3

       # Путь к relay логу
       relay-log = /var/log/mysql/mysql-relay-bin.log

       # Путь к bin логу на Мастере
       log_bin = /var/log/mysql/mysql-bin.log

       # название Вашей базы данных, которая будет реплицироваться
       binlog_do_db=socialdb"

## build master
#docker build --no-cache --build-arg sec="$R" -t sigma-social-db:latest .

## build slave1
#rm ./slave1/.my.cnf
#echo "$mycnf1" >> ./slave1/.my.cnf
#docker build --no-cache --build-arg sec="$R" -f ./slave1/Dockerfile -t sigma-social-db-replica-1:latest .

# build slave2
#rm ./slave1/.my.cnf
#echo "$mycnf2" >> ./slave1/.my.cnf
#docker build --no-cache --build-arg sec="$R" -f ./slave1/Dockerfile -t sigma-social-db-replica-2:latest .

#build proxysql
docker build --no-cache -f ./proxysql/Dockerfile -t sigma-social-proxysql:latest .