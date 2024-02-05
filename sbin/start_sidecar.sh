#!/bin/bash

# the local sbin paths are relative to the project root
SBIN=$(dirname "$(readlink -f "$0")")
SBIN="$(
  cd "$SBIN"
  pwd
)"
. $SBIN/utils/utils.sh
ROOT_DIR=$SBIN/..

WORKSPACE_DIR=$HOME/.spc/workspaces/active_workspace

PATHS_ENV=$WORKSPACE_DIR/.paths.env
SIDECAR_ENV=$WORKSPACE_DIR/.sidecar.env

reqdotenv "paths" $PATHS_ENV
reqdotenv "sidecar" $SIDECAR_ENV

spc up sidecar --validator --disseminator
