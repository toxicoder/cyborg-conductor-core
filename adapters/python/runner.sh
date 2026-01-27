#!/bin/bash

# Simple Bash entrypoint for the Python adapter
# This script calls the Python exec_wrapper.py with the provided arguments

# Make sure we have Python available
if ! command -v python3 &> /dev/null; then
    echo "Error: Python3 is required but not installed."
    exit 1
fi

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Execute the Python wrapper with all arguments
python3 "$SCRIPT_DIR/exec_wrapper.py" "$@"