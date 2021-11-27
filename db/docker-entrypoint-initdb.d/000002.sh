##!/bin/bash
set -e
source "../.msql"
mysql -u root -p"$R" -e "grant index on socialdb.* to social";
