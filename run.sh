#!/usr/bin/env bash
set -euo pipefail

cd $(dirname $0)

if command -v go &>/dev/null; then
    echo "Info: Go is installed."
else
    echo "Error: Go is not installed."
    exit 1
fi

echo "Info: Building and running binary."

go build .

./dotfiler

rm dotfiler
