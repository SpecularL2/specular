#!/bin/bash

# Converts a path to be relative to another directory.
relpath() {
  echo $(python3 -c "import os.path; print(os.path.relpath('$1', '$2'))")
}

# Requests a user to confirm that overwriting
# file ($1) is okay, if it exists.
guard_overwrite() {
  if test -f $1; then
    read -r -p "Overwrite $1 with a new file? [y/N] " response
    if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
      rm $1
    else
      exit 1
    fi
  fi
}

# Requires that a dotenv exists at path ($2).
reqdotenv() {
  if ! test -f "$2"; then
    echo "Expected $1 dotenv at $2 (not found)."
    exit
  fi
  echo "Using $1 dotenv: $2"
  . $2
}

generate_wallet() {
  wallet=$(cast wallet new)
  address=$(echo "$wallet" | awk '/Address/ { print $2 }')
  priv_key=$(echo "$wallet" | awk '/Private key/ { print $3 }')
  echo "$address"
  guard_overwrite $1
  echo $priv_key >$1
}

generate_jwt_secret() {
  openssl rand -hex 32
}
