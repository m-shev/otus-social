##!/bin/bash
set -e
R="secret"
SR="shard_secret"

shardList=("message_shard_1" "message_shard_2")

function generate {
  dockerEntry='##!/bin/bash
set -e
mysql -u root -p'"$R"' -e "create user '"'"$1"'"'@'"'"'%'"'"' identified with mysql_native_password by '"'"$SR"'"'";
mysql -u root -p'"$R"' -e "create database messagedb";
mysql -u root -p'"$R"' -e "grant create, drop, select, insert, update, delete, alter, references on messagedb.* to '"$1"'";'

  echo "$dockerEntry" > message/docker-entrypoint-initdb.d/000001.sh
}

for shard in "${shardList[@]}"; do
  generate "$shard"
  docker build --no-cache -f ./message/Dockerfile -t "$shard"-db:latest .
done

### build dialog db
#docker build --no-cache -f ./dialog/Dockerfile -t dialog-db:latest .
#
### build message db
#docker build --no-cache -f ./message/Dockerfile -t messge-db:latest .
