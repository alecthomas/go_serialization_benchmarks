#!/bin/sh
rm -f /tmp/temp_report.js
for f in `sed s://.*:: benchmarks.go |awk '/Name:/{print $2}'|sed s/[,]//g`; do 
   echo "package goserbench" > algname_test.go; 
   echo "const ALG_NAME = $f" >> algname_test.go; 
   go test -tags genreport -run TestGenerateReport
   lines=`wc -l ./report/data.js|awk '{print $1;}'`
   tail -$(($lines)) ./report/data.js|head -$(($lines-2))  >> /tmp/temp_report.js
   echo -e "\t}," >> /tmp/temp_report.js
done

lines=`wc -l /tmp/temp_report.js|awk '{print $1;}'`
echo "var data = [" > report/data.js
head -$(($lines-2)) /tmp/temp_report.js  >> ./report/data.js
echo -e "\t}" >> ./report/data.js
echo -n "];" >> ./report/data.js
rm -f /tmp/temp_report.js


echo "package goserbench" > algname_test.go; 
echo "const ALG_NAME = \"any\"" >> algname_test.go; 
   