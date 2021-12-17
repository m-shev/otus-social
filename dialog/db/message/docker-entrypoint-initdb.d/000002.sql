##!/bin/bash
set -e
mysql -u root -p"$R" -e "create user message_shard_1@% identified with mysql_native_password by secret";
mysql -u root -p"$R" -e "create database messagedb";
mysql -u root -p"$R" -e "grant create, drop, select, insert, update, delete, alter, references on messagedb.* to message_shard_1";
