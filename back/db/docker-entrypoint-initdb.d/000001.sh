##!/bin/bash
#set -e
mysqlp=123
upass='pass'

mysql -u "root" -p"$mysqlp" -e "create user 'social'@'%' identified by '"$upass"'";
#-e "create user 'social'@'%' identified $upass"
#-- #mysql -u root -p"$mysqlp"
#-- #create user social@% identified by 'pass';

#create user 'social'@'%' identified by 'pass';