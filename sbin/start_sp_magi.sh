#!/bin/bash

# the local sbin paths are relative to the project root
SBIN=$(dirname "$(readlink -f "$0")")
SBIN="$(
  cd "$SBIN"
  pwd
)"
. $SBIN/utils/utils.sh
ROOT_DIR=$SBIN/..
MAGI_WAIT=/tmp/sp_magi_started.lock

# Check that the all required dotenv files exists.
reqdotenv "paths" ".paths.env"
reqdotenv "sp_magi" ".sp_magi.env"

# Set sync flags.
SYNC_FLAGS=""
if [ $SYNC_MODE = "checkpoint" ]; then
  SYNC_FLAGS="--checkpoint-sync-url $CHECKPOINT_SYNC_URL --checkpoint-hash $CHECKPOINT_HASH"
fi

# Set devnet flags.
DEVNET_FLAGS=""
if [ "$DEVNET" = true ]; then
  echo "Enabling devnet mode."
  DEVNET_FLAGS="--devnet"
fi

# Set local sequencer flags.
SEQUENCER_FLAGS=""
if [ "$SEQUENCER" = true ]; then
  echo "Enabling local sequencer."
  SEQUENCER_FLAGS="
        --sequencer \
        --sequencer-max-safe-lag $SEQUENCER_MAX_SAFE_LAG \
        --sequencer-pk-file $SEQUENCER_PK_FILE"
fi

# TODO: use array for flags
FLAGS="
    --network $NETWORK \
    --l1-rpc-url $L1_RPC_URL \
    --l2-rpc-url $L2_RPC_URL \
    --sync-mode $SYNC_MODE \
    --l2-engine-url $L2_ENGINE_URL \
    --jwt-file $JWT_SECRET_PATH \
    --rpc-port $RPC_PORT \
    $SYNC_FLAGS $DEVNET_FLAGS $SEQUENCER_FLAGS $@"

echo "starting sp-magi with the following flags:"
echo "$FLAGS"
echo "Setting wait for file"
touch $MAGI_WAIT
$SP_MAGI_BIN $FLAGS

