# aoc2020

This repo contains my solution for aoc 2020, for details about aoc itself see [here](https://adventofcode.com/)

To run the code for a specific day, for example day 1, run `go run main.go d1`, commands take in custom input file locations, e.g. `go run main.go --input path/to/input/file`

Note, due to my stupidity most of this code will not handle CRLF. To use your input you will very likely need to convert the CRLF to LF first if you're on windows. 

```
Day by day solutions for the AOC 2020 puzzles

Usage:
  aoc2020 [command]

Available Commands:
  d1          Run solution for day 1
  d10         Run solution for day 10
  d11         Run solution for day 11
  d12         Run solution for day 12
  d13         Run solution for day 13
  d14         Run solution for day 14
  d15         Run solution for day 15
  d16         Run solution for day 16
  d17         Run solution for day 17
  d18         Run solution for day 18
  d19         Run solution for day 19
  d2          Run solution for day 2
  d20         Run solution for day 20
  d21         Run solution for day 21
  d3          Run solution for day 3
  d4          Run solution for day 4
  d5          Run solution for day 5
  d6          Run solution for day 6
  d7          Run solution for day 7
  d8          Run solution for day 8
  d9          Run solution for day 9
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.aoc2020.yaml)
  -h, --help            help for aoc2020
  -t, --toggle          Help message for toggle

Use "aoc2020 [command] --help" for more information about a command.
```