# Advent of Code 2025

My solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## Structure

```
aoc-2025/
├── cmd/
│   └── day01/          # Solution for Day 1
│       └── main.go
├── input/              # What I use to store input files (gitignored)
├── internal/
│   └── utils/
│       └── file.go
├── go.mod
└── README.md
```

## Prerequisites

- Go 1.25.4 or later

## Running Solutions

Each day's solution is in its own directory under `cmd/`. To run a specific day:

```bash
# Day 1
go run cmd/day01/main.go
```

## Input Files

Input files are stored in the `input/` directory and are gitignored. Place your puzzle inputs in this directory.