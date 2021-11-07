##!/bin/bash
set -e
source "../.msql"
echo "i work ------------------>"
mysql -u root -p"$R" -e "grant index on socialdb.* to social";
