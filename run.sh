#!/usr/bin/env bash
set -euo pipefail

cd $(dirname $0)

binary="dotbin"

if command -v go &>/dev/null; then
    echo "Info: Go is installed."
else
    echo "Error: Go is not installed."
    exit 1
fi

echo "Info: Building and running binary."

go build -o $binary ./src
trap 'rm -f "$binary"' EXIT

set +u
if [[ $1 == "dry-run" ]]; then
    ./$binary --dry-run
elif [[ $1 == "normal" ]]; then
    ./$binary
else
    echo "Usage: $0 [dry-run|normal]"
fi
set -u
