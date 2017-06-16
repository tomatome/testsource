#!/bin/sh
echo "1: copy /media/sf_VData/code4/source/src/test to local"
rm -rf test
rm -rf pkg/
cp /media/sf_VData/code4/source/src/test.go ./ -r
exitcode=$?
if [ $exitcode -eq 0 ];then
echo "copy done..."
else
echo "copy fail, exit"
exit $exitcode
fi
echo "2: go will build test"
go build -i -v -a test.go
exitcode=$?
if [ $exitcode -eq 0 ];then
echo "go build test done..."
else
echo "go build test fail, exit"
exit $exitcode
fi
echo "everything is done..."

