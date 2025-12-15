#!/bin/sh
set -e

# Function to replace HiGHS binary in dantzig dependency
replace_highs_binary() {
    # Search for dantzig dependency in common Livebook cache locations
    CACHE_DIRS="/home/livebook/.cache/livebook /root/.cache/livebook"
    
    for cache_dir in $CACHE_DIRS; do
        if [ -d "$cache_dir" ]; then
            # Find all dantzig installations
            find "$cache_dir" -type d -path "*/deps/dantzig/priv/solver/x86_64-linux-gnu" 2>/dev/null | while read -r solver_dir; do
                if [ -d "$solver_dir" ]; then
                    echo "Found dantzig at: $solver_dir"
                    cp /usr/local/bin/highs "$solver_dir/highs"
                    chmod +x "$solver_dir/highs"
                    echo "Replaced HiGHS binary in: $solver_dir"
                fi
            done
        fi
    done
}

# Replace HiGHS binary if dantzig is already installed
replace_highs_binary

# If running mix commands, ensure deps are compiled
if echo "$@" | grep -q "mix"; then
    cd /app
    mix deps.get 2>/dev/null || true
    mix deps.compile 2>/dev/null || true
fi

# Execute the command passed to the container
exec "$@"
