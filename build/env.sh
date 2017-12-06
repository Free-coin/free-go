#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"

# temporary free coin workspace
ethdir="$workspace/src/github.com/freecoin"
if [ ! -L "$ethdir/free-go" ]; then
    mkdir -p "$ethdir"
    cd "$ethdir"
    ln -s ../../../../../. free-go
    cd "$root"
fi

#

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$ethdir/free-go"
PWD="$ethdir/free-go"

# Launch the arguments with the configured environment.
exec "$@"
