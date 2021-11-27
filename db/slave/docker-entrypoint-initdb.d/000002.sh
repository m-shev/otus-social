##!/bin/bash
set -e
source "../.msql"

mysql -u root -p"$R" -e "INSTALL PLUGIN rpl_semi_sync_replica SONAME 'semisync_replica.so'";
mysql -u root -p"$R" -e "create user 'slave_user'@'%' IDENTIFIED WITH mysql_native_password BY '"$RPL"'";
mysql -u root -p"$R" -e "grant replication slave on *.* to slave_user;";