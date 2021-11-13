##!/bin/bash
set -e
source "../.msql"

mysql -u root -p"$R" -e "create database socialdb";
#mysql -u root -p"$R" socialdb < socialdb.sql;
#mysql -u root -p"$R" -e "CHANGE MASTER TO MASTER_HOST='db', MASTER_USER='slave_user', MASTER_PASSWORD='"$RPL"',  MASTER_LOG_FILE = 'mysql-bin.000001', MASTER_LOG_POS = 156;"
mysql -u root -p"$R" -e "CHANGE MASTER TO MASTER_HOST='db', MASTER_USER='slave_user', MASTER_PASSWORD='"$RPL"', MASTER_LOG_FILE = 'mysql-bin.000001', MASTER_LOG_POS = 4;"
