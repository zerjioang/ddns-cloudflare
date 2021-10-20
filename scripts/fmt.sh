#!/bin/bash

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "formatting source code with Go fmt"

packageName="."

#get all files excluding vendors
filelist=$(find ./ -name "*.go" | grep -v /vendor/)
toreplace="./"
toreplaceBy="/"
for file in $filelist
do
	# replace ./ to /
	file="${file/$toreplace/$toreplaceBy}"
	absolutePath=$packageName$file
	thisPackageName=$(dirname $absolutePath)
	echo "formatting package $thisPackageName"
	go fmt $thisPackageName
done

echo "code formatting done!"
