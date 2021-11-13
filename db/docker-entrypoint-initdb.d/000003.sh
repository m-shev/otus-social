##!/bin/bash
set -e
source "../.msql"
echo "-------------> create user 'slave_user'@'%'"
mysql -u root -p"$R" -e "create user 'slave_user'@'%' IDENTIFIED WITH mysql_native_password BY '"$RPL"'";
mysql -u root -p"$R" -e "grant replication slave on *.* to slave_user;";
echo "-------------> INSTALL PLUGIN rpl_semi_sync_source"
mysql -u root -p"$R" -e "INSTALL PLUGIN rpl_semi_sync_source SONAME 'semisync_source.so'";