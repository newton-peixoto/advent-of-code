#!/bin/sh
set -e

# Replace the x86_64 HiGHS binary with ARM64 version if it exists
if [ -f /usr/local/bin/highs ] && [ -d deps/dantzig/priv/solver/x86_64-linux-gnu ]; then
    cp /usr/local/bin/highs deps/dantzig/priv/solver/x86_64-linux-gnu/highs
    chmod +x deps/dantzig/priv/solver/x86_64-linux-gnu/highs
fi

# Execute the command passed to the container
exec "$@"
