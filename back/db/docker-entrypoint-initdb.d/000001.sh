##!/bin/bash
set -e
mysqlp=1234
upass='123qwe'

mysql -u root -p"$mysqlp" -e "create user 'social'@'%' identified by '"$upass"'";
mysql -u root -p"$mysqlp" -e "create database socialdb";
mysql -u root -p"$mysqlp" -e "grant create, insert, update, delete, alter on socialdb.* to social";

