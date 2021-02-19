#!/usr/bin/env bash

benchFuncs="${1}"
if [[ "${benchFuncs}0" == "0" ]] ;then {
    benchFuncs=".*"
}
fi

cat results.txt |  awk -F' ' '
BEGIN {
	print "benchmark                                     | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc"
	print "----------------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------"
}

/Benchmark/ {
	gsub(/ +/," ",$0)
	if ($7 != 0) {
	    printf "%-45s | %10d | %6d %s | %9d | %10d | %6.2f | %12d | %7.2f\n",$1,$2,$3,$4,$5,$7,$2*$3/1000000000,$2*$5/10000,$3/$7
	}else{
	    printf "%-45s | %10d | %6d %s | %9d | %10d | %6.2f | %12d | %7.2f\n",$1,$2,$3,$4,$5,$7,$2*$3/1000000000,$2*$5/10000,0
	}
	gsub(/(Unm|M)arshal/,"",$1)
	pname[$1]=$1
	for (i=2; i<=NF; i++) {
		if ($i ~ /[0-9]/) proto[$1,i]+=$i
		else proto[$1,i]=$i
	}
}

function arr_sort(arr,number) {
    i=0
    for (k in arr){
       number[i]=k
       i++
    }
    n = length(arr)
    for (i = 0; i < n; i++){
        for (j = 0; j < n; j++){
            if (number[j]+0 > number[i]+0){
                tmp = number[i]
                number[i] = number[j]
                number[j] = tmp
            }
        }
    }
}

END {
	print "\n"
	print "Totals:\n\n"
	print "benchmark                                | iter       | time/iter    | bytes/op  | allocs/op  | tt.sec | tt.kb        | ns/alloc"
	print "-----------------------------------------|------------|--------------|-----------|------------|--------|--------------|-----------"
	for (p in pname) {
		pr=pname[p]
	    arry[proto[pr,3]] = sprintf("%-40s | %10d | %6d %s | %9d | %10d | %6.2f | %12d | %7.2f",pr,proto[pr,2],proto[pr,3],proto[pr,4],proto[pr,5],proto[pr,7],
			proto[pr,2]*proto[pr,3]/1000000000,proto[pr,2]*proto[pr,5]/10000,proto[pr,3]/proto[pr,7])
	}
    arr_sort(arry,keys)
	for(i=0;i<length(keys);i++){
        print arry[keys[i]]
    }
}'
