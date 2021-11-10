##!/bin/bash
set -e
source "../.msql"
echo "-------------> INSTALL PLUGIN rpl_semi_sync_source"
mysql -u root -p"$R" -e "INSTALL PLUGIN rpl_semi_sync_source SONAME 'semisync_source.so'";