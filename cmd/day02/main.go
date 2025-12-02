package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/migvas/aoc-2025/internal/utils"
)

func main() {
	lines, err := utils.ReadLines("./input/d2.txt")

	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	part1(lines[0])
	part2(lines[0])
}

func part1(line string) {
	ranges := strings.Split(line, ",")
	invalidIdSum := 0

	for _, r := range ranges {
		parts := strings.Split(r, "-")

		// Convert the ranges to int so i can loop through them
		start, err := strconv.Atoi(parts[0])

		if err != nil {
			log.Fatalf("failed to parse string input: %v", err)
		}

		end, err := strconv.Atoi(parts[1])

		if err != nil {
			log.Fatalf("failed to parse string input: %v", err)
		}

		for i := start; i <= end; i++ {
			// Convert the int to string and compare the first half of the string with the second
			stringNum := strconv.Itoa(i)
			strLen := len(stringNum)

			// IDs with odd length can't be invalid
			if strLen%2 != 0 {
				continue
			}

			mid := strLen / 2
			if stringNum[:mid] == stringNum[mid:] {
				invalidIdSum += i
			}
		}
	}

	fmt.Printf("Part 1 - Invalid IDs sum: %d\n", invalidIdSum)
}

func part2(line string) {
	ranges := strings.Split(line, ",")
	invalidIdSum := 0

	for _, r := range ranges {
		parts := strings.Split(r, "-")

		// Convert the ranges to int so i can loop through them
		start, err := strconv.Atoi(parts[0])

		if err != nil {
			log.Fatalf("failed to parse string input: %v", err)
		}

		end, err := strconv.Atoi(parts[1])

		if err != nil {
			log.Fatalf("failed to parse string input: %v", err)
		}

		for i := start; i <= end; i++ {
			stringNum := strconv.Itoa(i)

			// Since the ID must be made of sequences the max sequence size
			// is half the size of the ID (for an ID made of 2 sequences)
			mid := len(stringNum) / 2
			// Check if the ID is made of sequences
			if isSequence(stringNum, stringNum[:mid]) {
				invalidIdSum += i
			}
		}
	}

	fmt.Printf("Part 2 - Invalid IDs sum: %d\n", invalidIdSum)
}

// Use a recursive function to check for a sequence inside the ID
func isSequence(id string, seq string) bool {
	// Stopping criteria, also stops checking for single digit IDs
	if len(seq) == 0 {
		return false
	}

	checkSeq := true

	for i := len(seq); i < len(id); i += len(seq) {
		if i+len(seq) > len(id) || id[i:i+len(seq)] != seq {
			checkSeq = false
			break
		}
	}

	if checkSeq {
		return true
	}

	// If that sequence doesn't repeat itself then we check for a sequence
	// with one character less
	return isSequence(id, seq[:len(seq)-1])
}
