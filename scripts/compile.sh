#!/bin/bash

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..
ls -alh

echo "compiling, please wait"
CGO_ENABLED=0 go build \
  -ldflags "-w -s -extldflags=-static" \
  -a \
  -tags osusergo,netgo \
  -trimpath \
  -o cfagent && \

echo "compilation finished"
ls -alh cfagent

# upx --brute cfagent

# go releaser
# go install github.com/goreleaser/goreleaser@latest
# goreleaser init
# goreleaser release --snapshot --rm-dist
# goreleaser check
# goreleaser build --single-target
# export GITHUB_TOKEN="YOUR_GH_TOKEN"
# git tag -a v0.1.0 -m "First release"
  #git push origin v0.1.0