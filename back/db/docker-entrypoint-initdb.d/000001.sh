##!/bin/bash
set -e
mysqlp=1234
upass='123qwe'

mysql -u root -p"$mysqlp" -e "create user 'social'@'%' identified by '"$upass"'";
mysql -u root -p"$mysqlp" -e "create database socialdb";
mysql -u root -p"$mysqlp" -e "grant create, drop, select, insert, update, delete, alter, references on socialdb.* to social";

