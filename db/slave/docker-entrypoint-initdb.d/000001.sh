##!/bin/bash
set -e
source "../.msql"

mysql -u root -p"$R" -e "create user 'social'@'%' IDENTIFIED BY '"$S"'";
mysql -u root -p"$R" -e "create database socialdb";
mysql -u root -p"$R" -e "grant create, drop, select, insert, update, delete, alter, references on socialdb.* to social";
mysql -u root -p"$R" socialdb < socialdb.sql;
mysql -u root -p"$R" -e "CHANGE MASTER TO MASTER_HOST='db', MASTER_USER='slave_user', MASTER_PASSWORD='"$RPL"', MASTER_LOG_FILE = 'mysql-bin.000005', MASTER_LOG_POS = 156;"
