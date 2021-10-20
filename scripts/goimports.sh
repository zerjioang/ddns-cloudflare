#!/bin/bash

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "checking imports in source code with goimports"

echo "checking if goimports is installed in $GOPATH"
if [[ ! -f ${GOPATH}/bin/goimports ]]; then
	#statements
	echo "goimports not found. Downloading via go get"
	go install golang.org/x/tools/cmd/goimports
fi

#get all files excluding vendors
filelist=$(find . -type f -name "*.go" | grep -vendor)
for file in ${filelist}
do
	echo "goimports check in file $file"
	${GOPATH}/bin/goimports -v -w ${file}
done

echo "code formatting done!"
