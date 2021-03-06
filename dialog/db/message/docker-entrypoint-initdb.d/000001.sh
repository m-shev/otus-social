##!/bin/bash
set -e
mysql -u root -psecret -e "create user 'message_shard_2'@'%' identified with mysql_native_password by 'shard_secret'";
mysql -u root -psecret -e "create database messagedb";
mysql -u root -psecret -e "grant create, drop, select, insert, update, delete, alter, references on messagedb.* to message_shard_2";
