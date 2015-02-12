Added totals by protocol (marshal+unmarshal) and total time+bytes and time/alloc to allow better comparisons (using ./stats)

```
benchmark                                      iter    time/iter   bytes/op     allocs/op  tt.time   tt.bytes       time/alloc
BenchmarkVmihailencoMsgpackMarshal          1000000   2735 ns/op   355 B/op   5 allocs/op   2.73 s   35500 KB  547.00 ns/alloc
BenchmarkVmihailencoMsgpackUnmarshal         500000   3220 ns/op   352 B/op   9 allocs/op   1.61 s   17600 KB  357.78 ns/alloc
BenchmarkJsonMarshal                         500000   5205 ns/op   590 B/op   7 allocs/op   2.60 s   29500 KB  743.57 ns/alloc
BenchmarkJsonUnmarshal                       200000   8012 ns/op   468 B/op   7 allocs/op   1.60 s    9360 KB 1144.57 ns/alloc
BenchmarkBsonMarshal                        1000000   2950 ns/op   488 B/op  13 allocs/op   2.95 s   48800 KB  226.92 ns/alloc
BenchmarkBsonUnmarshal                       500000   3456 ns/op   281 B/op  10 allocs/op   1.73 s   14050 KB  345.60 ns/alloc
BenchmarkVitessBsonMarshal                  1000000   2056 ns/op  1169 B/op   4 allocs/op   2.06 s  116900 KB  514.00 ns/alloc
BenchmarkVitessBsonUnmarshal                1000000   1114 ns/op   227 B/op   4 allocs/op   1.11 s   22700 KB  278.50 ns/alloc
BenchmarkGobMarshal                          200000   9812 ns/op  1665 B/op  25 allocs/op   1.96 s   33300 KB  392.48 ns/alloc
BenchmarkGobUnmarshal                         50000  72331 ns/op 19298 B/op 366 allocs/op   3.62 s   96490 KB  197.63 ns/alloc
BenchmarkXdrMarshal                          500000   3904 ns/op   519 B/op  15 allocs/op   1.95 s   25950 KB  260.27 ns/alloc
BenchmarkXdrUnmarshal                       1000000   2940 ns/op   274 B/op   9 allocs/op   2.94 s   27400 KB  326.67 ns/alloc
BenchmarkUgorjiCodecMsgpackMarshal           500000   4930 ns/op  1922 B/op  10 allocs/op   2.46 s   96100 KB  493.00 ns/alloc
BenchmarkUgorjiCodecMsgpackUnmarshal         500000   4787 ns/op  1857 B/op  11 allocs/op   2.39 s   92850 KB  435.18 ns/alloc
BenchmarkUgorjiCodecBincMarshal              500000   4967 ns/op  1952 B/op  10 allocs/op   2.48 s   97600 KB  496.70 ns/alloc
BenchmarkUgorjiCodecBincUnmarshal            500000   5311 ns/op  2020 B/op  14 allocs/op   2.66 s  101000 KB  379.36 ns/alloc
BenchmarkSerealMarshal                       500000   7344 ns/op  1373 B/op  24 allocs/op   3.67 s   68650 KB  306.00 ns/alloc
BenchmarkSerealUnmarshal                     500000   5804 ns/op   999 B/op  24 allocs/op   2.90 s   49950 KB  241.83 ns/alloc
BenchmarkBinaryMarshal                      1000000   2864 ns/op   406 B/op  12 allocs/op   2.86 s   40600 KB  238.67 ns/alloc
BenchmarkBinaryUnmarshal                     500000   3443 ns/op   434 B/op  17 allocs/op   1.72 s   21700 KB  202.53 ns/alloc
BenchmarkMsgpMarshal                        5000000    543 ns/op   144 B/op   1 allocs/op   2.71 s   72000 KB  543.00 ns/alloc
BenchmarkMsgpUnmarshal                      2000000    715 ns/op   113 B/op   3 allocs/op   1.43 s   22600 KB  238.33 ns/alloc
BenchmarkGoprotobufMarshal                  2000000   1178 ns/op   314 B/op   3 allocs/op   2.36 s   62800 KB  392.67 ns/alloc
BenchmarkGoprotobufUnmarshal                1000000   1615 ns/op   440 B/op   9 allocs/op   1.61 s   44000 KB  179.44 ns/alloc
BenchmarkGogoprotobufMarshal                5000000    285 ns/op    64 B/op   1 allocs/op   1.43 s   32000 KB  285.00 ns/alloc
BenchmarkGogoprotobufUnmarshal              5000000    412 ns/op   113 B/op   3 allocs/op   2.06 s   56500 KB  137.33 ns/alloc
BenchmarkProtobufMarshal                    1000000   1705 ns/op   214 B/op   6 allocs/op   1.71 s   21400 KB  284.17 ns/alloc
BenchmarkProtobufUnmarshal                  1000000   1701 ns/op   193 B/op   7 allocs/op   1.70 s   19300 KB  243.00 ns/alloc
---
totals:
BenchmarkBson                               1500000   6406 ns/op   769 B/op  23    9.61 s  115350 KB  278.52 ns/alloc
BenchmarkMsgp                               7000000   1258 ns/op   257 B/op   4    8.81 s  179900 KB  314.50 ns/alloc
BenchmarkVmihailencoMsgpack                 1500000   5955 ns/op   707 B/op  14    8.93 s  106050 KB  425.36 ns/alloc
BenchmarkVitessBson                         2000000   3170 ns/op  1396 B/op   8    6.34 s  279200 KB  396.25 ns/alloc
BenchmarkGogoprotobuf                      10000000    697 ns/op   177 B/op   4    6.97 s  177000 KB  174.25 ns/alloc
BenchmarkProtobuf                           2000000   3406 ns/op   407 B/op  13    6.81 s   81400 KB  262.00 ns/alloc
BenchmarkGoprotobuf                         3000000   2793 ns/op   754 B/op  12    8.38 s  226200 KB  232.75 ns/alloc
BenchmarkJson                                700000  13217 ns/op  1058 B/op  14    9.25 s   74060 KB  944.07 ns/alloc
BenchmarkUgorjiCodecMsgpack                 1000000   9717 ns/op  3779 B/op  21    9.72 s  377900 KB  462.71 ns/alloc
BenchmarkGob                                 250000  82143 ns/op 20963 B/op 391   20.54 s  524075 KB  210.08 ns/alloc
BenchmarkUgorjiCodecBinc                    1000000  10278 ns/op  3972 B/op  24   10.28 s  397200 KB  428.25 ns/alloc
BenchmarkBinary                             1500000   6307 ns/op   840 B/op  29    9.46 s  126000 KB  217.48 ns/alloc
BenchmarkSereal                             1000000  13148 ns/op  2372 B/op  48   13.15 s  237200 KB  273.92 ns/alloc
BenchmarkXdr                                1500000   6844 ns/op   793 B/op  24   10.27 s  118950 KB  285.17 ns/alloc
```

Considering only totals, sorting by:

1. total time
```
BenchmarkVitessBson                         2000000   3170 ns/op  1396 B/op   8    6.34 s  279200 KB  396.25 ns/alloc
BenchmarkProtobuf                           2000000   3406 ns/op   407 B/op  13    6.81 s   81400 KB  262.00 ns/alloc
BenchmarkGogoprotobuf                      10000000    697 ns/op   177 B/op   4    6.97 s  177000 KB  174.25 ns/alloc
BenchmarkGoprotobuf                         3000000   2793 ns/op   754 B/op  12    8.38 s  226200 KB  232.75 ns/alloc
BenchmarkMsgp                               7000000   1258 ns/op   257 B/op   4    8.81 s  179900 KB  314.50 ns/alloc
BenchmarkVmihailencoMsgpack                 1500000   5955 ns/op   707 B/op  14    8.93 s  106050 KB  425.36 ns/alloc
BenchmarkJson                                700000  13217 ns/op  1058 B/op  14    9.25 s   74060 KB  944.07 ns/alloc
BenchmarkBinary                             1500000   6307 ns/op   840 B/op  29    9.46 s  126000 KB  217.48 ns/alloc
BenchmarkBson                               1500000   6406 ns/op   769 B/op  23    9.61 s  115350 KB  278.52 ns/alloc
BenchmarkUgorjiCodecMsgpack                 1000000   9717 ns/op  3779 B/op  21    9.72 s  377900 KB  462.71 ns/alloc
BenchmarkXdr                                1500000   6844 ns/op   793 B/op  24   10.27 s  118950 KB  285.17 ns/alloc
BenchmarkUgorjiCodecBinc                    1000000  10278 ns/op  3972 B/op  24   10.28 s  397200 KB  428.25 ns/alloc
BenchmarkSereal                             1000000  13148 ns/op  2372 B/op  48   13.15 s  237200 KB  273.92 ns/alloc
BenchmarkGob                                 250000  82143 ns/op 20963 B/op 391   20.54 s  524075 KB  210.08 ns/alloc
```

2. total memory
```
BenchmarkJson                                700000  13217 ns/op  1058 B/op  14    9.25 s   74060 KB  944.07 ns/alloc
BenchmarkProtobuf                           2000000   3406 ns/op   407 B/op  13    6.81 s   81400 KB  262.00 ns/alloc
BenchmarkVmihailencoMsgpack                 1500000   5955 ns/op   707 B/op  14    8.93 s  106050 KB  425.36 ns/alloc
BenchmarkBson                               1500000   6406 ns/op   769 B/op  23    9.61 s  115350 KB  278.52 ns/alloc
BenchmarkXdr                                1500000   6844 ns/op   793 B/op  24   10.27 s  118950 KB  285.17 ns/alloc
BenchmarkBinary                             1500000   6307 ns/op   840 B/op  29    9.46 s  126000 KB  217.48 ns/alloc
BenchmarkGogoprotobuf                      10000000    697 ns/op   177 B/op   4    6.97 s  177000 KB  174.25 ns/alloc
BenchmarkMsgp                               7000000   1258 ns/op   257 B/op   4    8.81 s  179900 KB  314.50 ns/alloc
BenchmarkGoprotobuf                         3000000   2793 ns/op   754 B/op  12    8.38 s  226200 KB  232.75 ns/alloc
BenchmarkSereal                             1000000  13148 ns/op  2372 B/op  48   13.15 s  237200 KB  273.92 ns/alloc
BenchmarkVitessBson                         2000000   3170 ns/op  1396 B/op   8    6.34 s  279200 KB  396.25 ns/alloc
BenchmarkUgorjiCodecMsgpack                 1000000   9717 ns/op  3779 B/op  21    9.72 s  377900 KB  462.71 ns/alloc
BenchmarkUgorjiCodecBinc                    1000000  10278 ns/op  3972 B/op  24   10.28 s  397200 KB  428.25 ns/alloc
BenchmarkGob                                 250000  82143 ns/op 20963 B/op 391   20.54 s  524075 KB  210.08 ns/alloc
```

3. speed per alloc (not sure if meaningful)
```
BenchmarkGogoprotobuf                      10000000    697 ns/op   177 B/op   4    6.97 s  177000 KB  174.25 ns/alloc
BenchmarkGob                                 250000  82143 ns/op 20963 B/op 391   20.54 s  524075 KB  210.08 ns/alloc
BenchmarkBinary                             1500000   6307 ns/op   840 B/op  29    9.46 s  126000 KB  217.48 ns/alloc
BenchmarkGoprotobuf                         3000000   2793 ns/op   754 B/op  12    8.38 s  226200 KB  232.75 ns/alloc
BenchmarkProtobuf                           2000000   3406 ns/op   407 B/op  13    6.81 s   81400 KB  262.00 ns/alloc
BenchmarkSereal                             1000000  13148 ns/op  2372 B/op  48   13.15 s  237200 KB  273.92 ns/alloc
BenchmarkBson                               1500000   6406 ns/op   769 B/op  23    9.61 s  115350 KB  278.52 ns/alloc
BenchmarkXdr                                1500000   6844 ns/op   793 B/op  24   10.27 s  118950 KB  285.17 ns/alloc
BenchmarkMsgp                               7000000   1258 ns/op   257 B/op   4    8.81 s  179900 KB  314.50 ns/alloc
BenchmarkVitessBson                         2000000   3170 ns/op  1396 B/op   8    6.34 s  279200 KB  396.25 ns/alloc
BenchmarkVmihailencoMsgpack                 1500000   5955 ns/op   707 B/op  14    8.93 s  106050 KB  425.36 ns/alloc
BenchmarkUgorjiCodecBinc                    1000000  10278 ns/op  3972 B/op  24   10.28 s  397200 KB  428.25 ns/alloc
BenchmarkUgorjiCodecMsgpack                 1000000   9717 ns/op  3779 B/op  21    9.72 s  377900 KB  462.71 ns/alloc
BenchmarkJson                                700000  13217 ns/op  1058 B/op  14    9.25 s   74060 KB  944.07 ns/alloc
```

