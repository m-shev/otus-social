##!/bin/bash
set -e
R="secret"
mysql -u root -p"$R" -e "create user 'dialog'@'%' identified with mysql_native_password by '"$R"'";
mysql -u root -p"$R" -e "create database dialogdb";
mysql -u root -p"$R" -e "grant create, drop, select, insert, update, delete, alter, references on dialogdb.* to dialog";