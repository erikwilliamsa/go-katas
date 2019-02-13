# File Parsing

Created 3 different methods of parsing files and getting word counts as well as alpha numeric character counts.

### BlockingAllFileStats

BlockingAllFileStats loops through all files in a given directory ( ignoring sub directories ).  For each file it will read the lines in the file and count the words then count the alpha numeric characters.


### ConcurrentFileStats

ConcurrentFileStats does the same as BlockingAllFileStats except that for each file it kicks off a go routine to gather the stats.

### PipedGetAllFileStats

PipedGetAllFileStats creates a data pipeline using multiple channels and go routines for each step. In the end, channels are merged back in to create a single output.

1. For each file in a directory, create 
    - go routine for counting words in a line with an channel that outputs an `int`
    - go routine fo counting characters in a line with a channel that outputs a map containing the alpha numeric characters and `int` count of number of times the character was found in that string.

2. For each line of a file  send it over the word count and character count channels
3. Concurrently count words and characters for a given line and output the results over the channel that a collector is listening to.
4. Collector will accumulate the stats for a particular file in that pipeline.
5. Merge will get total up the counts for characters (merging ht maps as well), as well as combine a list of individual file stats and write it to a final output channel.


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