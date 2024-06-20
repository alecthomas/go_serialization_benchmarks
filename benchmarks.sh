#!/bin/sh
for f in `sed s://.*:: benchmarks.go |awk '/Name:/{print $2}'|sed s/[,]//g`; do 
   echo "package goserbench" > algname_test.go; 
   echo "const ALG_NAME = $f" >> algname_test.go; 
   go test -bench=. -benchmem| grep BenchmarkSerializers| sed s:BenchmarkSerializers/::; 
done
echo "package goserbench" > algname_test.go; 
echo "const ALG_NAME = \"any\"" >> algname_test.go; 



