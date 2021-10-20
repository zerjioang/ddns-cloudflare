#!/bin/bash

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "formatting source code with gofmt (and simplification)"

packageName="github.com/zerjioang/cf-agent"

#get all files excluding vendors
filelist=$(find ./ -name "*.go" | grep -vendor)
toreplace="./"
toreplaceBy="/"
for file in ${filelist}
do
	echo "formatting file $file"
	gofmt -s -w ${file}
done

echo "code formatting done!"
