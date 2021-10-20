#!/bin/bash

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Checking code quality with linters..."

# go fmt
./scripts/fmt.sh
# go fmt simplify
./scripts/fmt_and_simplify.sh
# go imports
./scripts/goimports.sh
# go vet
./scripts/govet.sh
# add license header to files
# ./scripts/license_header.sh

echo "qa scripts finished"
