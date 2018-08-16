#!/usr/bin/env bash

benchFuncs="${1}"
if [[ "${benchFuncs}0" == "0" ]] ;then {
    benchFuncs=".*"
}
fi

go test -bench="${benchFuncs}" ./ |  awk -F' ' '
BEGIN {
	print "benchmark                                      iter    time/iter   bytes/op     allocs/op  tt.time   tt.bytes       time/alloc"
	print "---------                                      ----    ---------   --------     ---------  -------   --------       ----------"
}

/Benchmark/ {
	gsub(/ +/," ",$0)
	if ($7 != 0) {
	    printf "%-40s %10d %6d %s %5d %s %3d %s %6.2f s %7d KB %7.2f ns/alloc\n",$1,$2,$3,$4,$5,$6,$7,$8,$2*$3/1000000000,$2*$5/10000,$3/$7
	}else{
	    printf "%-40s %10d %6d %s %5d %s %3d %s %6.2f s %7d KB %7.2f ns/alloc\n",$1,$2,$3,$4,$5,$6,$7,$8,$2*$3/1000000000,$2*$5/10000,0
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
	print "---\ntotals:"
	for (p in pname) {
		pr=pname[p]
	    arry[proto[pr,3]] = sprintf("%-40s %10d %6d %s %5d %s %3d %s %6.2f s %7d KB %7.2f ns/alloc",pr,proto[pr,2],proto[pr,3],proto[pr,4],proto[pr,5],proto[pr,6],proto[pr,7],
			proto[pr,8],proto[pr,2]*proto[pr,3]/1000000000,proto[pr,2]*proto[pr,5]/10000,proto[pr,3]/proto[pr,7])
	}
    arr_sort(arry,keys)
	for(i=0;i<length(keys);i++){
        print arry[keys[i]]
    }
}'