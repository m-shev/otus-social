##!/bin/bash
set -e
source "../.msql"

mysql -u root -p"$R" -e "create user 'social'@'%' identified by '"$S"'";
mysql -u root -p"$R" -e "create database socialdb";
mysql -u root -p"$R" -e "grant create, drop, select, insert, update, delete, alter, references on socialdb.* to social";

