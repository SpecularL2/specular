#!/bin/bash

# TODO: can we get rid of this somehow?
# currently the local sbin paths are relative to the project root
SBIN=$(dirname "$(readlink -f "$0")")
SBIN="`cd "$SBIN"; pwd`"
ROOT_DIR=$SBIN/..

# Check that the all required dotenv files exists.
CONFIGURE_ENV=".configure.env"
if ! test -f $CONFIGURE_ENV; then
    echo "Expected dotenv at $CONFIGURE_ENV (does not exist)."
    exit
fi
echo "Using configure dotenv: $CONFIGURE_ENV"
. $CONFIGURE_ENV

SIDECAR_ENV=".sidecar.env"
if ! test -f $SIDECAR_ENV; then
    echo "Expected dotenv at $SIDECAR_ENV (does not exist)."
    exit
fi
echo "Using sidecar dotenv: $SIDECAR_ENV"
. $SIDECAR_ENV

FLAGS=(
    "--l1.endpoint $L1_ENDPOINT"
    "--l2.endpoint $L2_ENDPOINT"
    "--protocol.rollup-cfg-path $ROLLUP_CFG_PATH"
    "--protocol.rollup-addr $ROLLUP_ADDR"
)

# Set disseminator flags.
if [ "$DISSEMINATOR" = true ] ; then
    echo "Enabling disseminator."
    DISSEMINATOR_PRIV_KEY=`cat "$DISSEMINATOR_PK_PATH"`
    FLAGS+=(
        "--disseminator"
        "--disseminator.private-key $DISSEMINATOR_PRIV_KEY"
        "--disseminator.sub-safety-margin $DISSEMINATOR_SUB_SAFETY_MARGIN"
        "--disseminator.target-batch-size $DISSEMINATOR_TARGET_BATCH_SIZE"
    )
fi
# Set validator flags.
if [ "$VALIDATOR" = true ] ; then
    echo "Enabling validator."
    VALIDATOR_PRIV_KEY=`cat "$VALIDATOR_PK_PATH"`
    FLAGS+=(
        "--validator"
        "--validator.private-key $VALIDATOR_PRIV_KEY"
    )
fi

echo "starting sidecar with the following flags:"
echo "${FLAGS[@]}"
$SIDECAR_BIN ${FLAGS[@]}
