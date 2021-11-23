##!/bin/bash
set -e
source "../.msql"

mysql -u root -p"$R" -e "INSTALL PLUGIN rpl_semi_sync_replica SONAME 'semisync_replica.so'";