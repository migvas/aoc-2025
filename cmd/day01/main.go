package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines, err := readLines("./input/d1.txt")

	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	part1(lines)
	part2(lines)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func part1(rotations []string) {
	zeroes := 0
	dialPos := 50

	for _, l := range rotations {
		direction := string(l[0])
		turns, err := strconv.Atoi(l[1:])
		if err != nil {
			log.Fatalf("failed to parse string input: %v", err)
		}

		if direction == "R" {
			dialPos += turns
		} else {
			dialPos -= turns
		}

		dialPos = dialPos % 100

		if dialPos < 0 {
			dialPos += 100
		}

		if dialPos == 0 {
			zeroes++
		}
	}

	fmt.Printf("Part 1 - Number of zero positions: %d\n", zeroes)
}

func part2(rotations []string) {
	zeroes := 0
	dialPos := 50

	for _, l := range rotations {
		direction := string(l[0])
		turns, err := strconv.Atoi(l[1:])
		startPos := dialPos

		if err != nil {
			log.Fatalf("failed to parse string input: %v", err)
		}

		if direction == "R" {
			dialPos += turns
		} else {
			dialPos -= turns
		}

		fullRotations := dialPos / 100

		if fullRotations < 0 {
			fullRotations = -fullRotations
		}
		if dialPos < 0 && startPos != 0 {
			fullRotations += 1
		}

		dialPos = dialPos % 100

		if dialPos < 0 {
			dialPos += 100
		}

		if fullRotations > 0 {
			zeroes += fullRotations
		} else if dialPos == 0 {
			zeroes++
		}
	}

	fmt.Printf("Part 2 - Number of zero positions: %d\n", zeroes)
}
