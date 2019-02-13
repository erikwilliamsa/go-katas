# File Parsing

Created 3 different methods of parsing files and getting word counts as well as alpha numeric character counts.

### BlockingAllFileStats

BlockingAllFileStats loops through all files in a given directory ( ignoring sub directories ).  For each file it will read the lines in the file and count the words then count the alpha numeric characters.


### ConcurrentFileStats

ConcurrentFileStats does the same as BlockingAllFileStats except that for each file it kicks off a go routine to gather the stats.

### PipedGetAllFileStats

PipedGetAllFileStats creates a data pipe line using multiple channels and go routines for each step.


```
                                              .- WordCount (number)        -> 
                                             /                                \__ Collector (file stats)->     
                      .-file 1 (each line)-><                                 /                            \
                     /                       \.- CharCount (map{string}num)->                               \_ merge -> All Stats
                    /                                                                                       /
dir: (each file)-->:                          .- WordCount (number)        ->                              /
                    \                        /                                 \__ Collector (file stats)-> 
                     \.-file 2 (each line)-><                                  /
                                             \.- CharCount (map{string}num)-> 
```