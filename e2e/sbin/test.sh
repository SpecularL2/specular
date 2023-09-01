#!/bin/bash

# Configure variables
SBIN_DIR=`dirname $0`
SBIN_DIR="`cd "$SBIN_DIR"; pwd`"
set -o allexport
source $SBIN_DIR/configure.sh
set +o allexport

# Spin up L1 node
cd $CONTRACTS_DIR
anvil --block-time 1 > $PROJECT_LOG_DIR/l1.log 2>&1 &
ANVIL_PID=$!
npx hardhat deploy --network localhost

# Spin up L2 node
cd $PROJECT_DATA_DIR
$SBIN_DIR/sequencer.sh > $PROJECT_LOG_DIR/l2.log 2>&1 &
L2GETH_PID=$!

# Wait for nodes
sleep 10
$SBIN_DIR/wait-for-it.sh -t 60 $HOST:$L1_WS_PORT
$SBIN_DIR/wait-for-it.sh -t 60 $HOST:$L2_HTTP_PORT

# Run testing script
cd $CONTRACTS_DIR
npx ts-node scripts/testing.ts
exit_0 = $?
npx hardhat run scripts/bridge/standard_bridge_deposit_eth.ts
exit_1 = $?
npx hardhat run scripts/bridge/standard_bridge_withdraw_eth.ts
exit_2 = $?
npx hardhat run scripts/bridge/standard_bridge_erc20.ts
exit_3 = $?

# Kill nodes
disown $L2GETH_PID
disown $ANVIL_PID
kill $L2GETH_PID
kill $ANVIL_PID

# Clean up
$SBIN_DIR/clean.sh

# Exit with 0 iff all test return status 0
! (( $exit_0 || $exit_1 || $exit_2 || $exit_3 ))
