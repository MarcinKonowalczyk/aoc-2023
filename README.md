# aoc-2023

[Advent of Code](https://adventofcode.com) 2023 in Go.


```
go run . -day 1 -part 1 -filename ./data/full/day01.txt -v
```

or

```
go build .
aoc2023 -day 1 -part 1 -filename ./data/full/day01.txt -v
```


## test

```
go test aoc2023/utils -v && \
go test aoc2023/dijkstra -v
```

or

```
~/go/bin/gotest aoc2023/utils -v && \
~/go/bin/gotest aoc2023/dijkstra -v  
```