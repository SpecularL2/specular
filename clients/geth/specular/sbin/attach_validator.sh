#!/bin/bash
SBIN=`dirname $0`
SBIN="`cd "$SBIN"; pwd`"
. $SBIN/configure.sh
cd $DATA_DIR
../build/bin/geth --datadir ./data_validator attach http://localhost:4018
