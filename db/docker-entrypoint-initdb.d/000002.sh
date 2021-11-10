##!/bin/bash
set -e
source "../.msql"
echo "-------------> grant index on socialdb.* to social"
mysql -u root -p"$R" -e "grant index on socialdb.* to social";
