#!/usr/bin/env bash
set -euo pipefail

cd $(dirname $0)

binary="dotbin"

set +u
if [[ $1 == "dry-run" ]]; then
    dry_run="yes"
elif [[ $1 == "normal" ]]; then
    dry_run="no"
else
    echo "Usage: $0 [dry-run|normal]"
    exit 1
fi
set -u

if command -v go &>/dev/null; then
    echo "Info: Go is installed"
else
    echo "Error: Go is not installed"
    exit 1
fi

echo "Info: Building and running binary"

go build -o $binary ./src
trap 'rm -f "$binary"' EXIT

if [[ $dry_run == "yes" ]]; then
    ./$binary --dry-run
elif [[ $dry_run == "no" ]]; then
    ./$binary
fi
