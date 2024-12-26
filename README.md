# AdventOfCode24
Advent of Code 2024 problems

- [AdventOfCode24](#adventofcode24)
  - [Day 01](#day-01)


## Day 01
For [day one](https://adventofcode.com/2024/day/1), I've added a Cobra CLI app which can analyse the required lists, as well as generate the left and right lists.

You can evaluate the CLI app like any other Cobra app via
```bash
go run . -h
```

The core commands are
```bash
# Evaluate left and right list for the total distance
go run . listLeft.txt listRight.txt

# Generate left and right lists. 
# It supports additional config flags to generate varying number of
# items and ID ranges
go run . gen listLeft.txt listRight.txt
```