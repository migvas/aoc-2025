package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges, ids, err := readLines("./input/d5.txt")

	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	part1(ranges, ids)
	part2(ranges)
}

func readLines(path string) ([]string, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Get ranges, when a blank line is found it stops
	var ranges []string
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		ranges = append(ranges, scanner.Text())
	}

	// After the blank line get the ids
	var ids []string
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		ids = append(ids, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return ranges, ids, nil
}

func part1(ranges []string, ids []string) {
	freshIDs := 0

	var intRanges [][]int

	// Build an array of arrays with the int values of the start and end of each range
	for _, r := range ranges {
		rangePos := strings.Split(r, "-")
		rangeStart, err := strconv.Atoi(rangePos[0])

		if err != nil {
			log.Fatalf("failed to parse string range start: %v", err)
		}

		rangeEnd, err := strconv.Atoi(rangePos[1])

		if err != nil {
			log.Fatalf("failed to parse string range end: %v", err)
		}

		intRanges = append(intRanges, []int{rangeStart, rangeEnd})
	}

	for _, id := range ids {
		// To check if the id is fresh we need to convert it to an int
		intID, err := strconv.Atoi(id)

		if err != nil {
			log.Fatalf("failed to parse string id: %v", err)
		}

		// Then we check each range to see if the id is inside it
		// If we find a match, we mark the id as fresh and break the loop
		for _, r := range intRanges {
			if intID >= r[0] && intID <= r[1] {
				freshIDs++
				break
			}
		}
	}

	fmt.Printf("Part 1 - Number of fresh IDs: %d\n", freshIDs)
}

func part2(ranges []string) {
	freshIDs := 0

	var intRanges [][2]int

	// Build an array of arrays with the int values of the start and end of each range
	for _, r := range ranges {
		rangePos := strings.Split(r, "-")
		rangeStart, err := strconv.Atoi(rangePos[0])

		if err != nil {
			log.Fatalf("failed to parse string range start: %v", err)
		}

		rangeEnd, err := strconv.Atoi(rangePos[1])

		if err != nil {
			log.Fatalf("failed to parse string range end: %v", err)
		}

		intRanges = append(intRanges, [2]int{rangeStart, rangeEnd})
	}

	// Create an infinite loop to merge the ranges
	for {
		// This map will hold the record of which ranges have been merged
		mergedRanges := make(map[[2]int]bool)
		// This slice will be used to store the next iteration of ranges to merge
		next := make([][2]int, 0, len(intRanges))

		for i := 0; i < len(intRanges); i++ {
			// For each range we will try to merge it with the other ranges
			currentRange := intRanges[i]
			// If it was already merged in this iteration we skip it
			if _, ok := mergedRanges[currentRange]; ok {
				continue
			}
			for j := i + 1; j < len(intRanges); j++ {
				// This is how we evalute if 2 ranges should be merged
				// if the conditions are met we append a merged range to the
				// next slice and mark both the currentRange and the target range as merged
				if currentRange[0] <= intRanges[j][0] {
					if currentRange[1] >= intRanges[j][1] {
						next = append(next, currentRange)

						mergedRanges[currentRange] = true
						mergedRanges[intRanges[j]] = true
					} else if currentRange[1] >= intRanges[j][0] {
						next = append(next, [2]int{currentRange[0], intRanges[j][1]})
						mergedRanges[currentRange] = true
						mergedRanges[intRanges[j]] = true
					}
				} else {
					if currentRange[1] <= intRanges[j][1] {
						next = append(next, intRanges[j])
						mergedRanges[currentRange] = true
						mergedRanges[intRanges[j]] = true
					} else if currentRange[0] <= intRanges[j][1] {
						next = append(next, [2]int{intRanges[j][0], currentRange[1]})
						mergedRanges[currentRange] = true
						mergedRanges[intRanges[j]] = true
					}
				}
			}

			// If after evaluating every range the current range wasn't merged then
			// we append it to the next slice as it is
			if _, ok := mergedRanges[currentRange]; !ok {
				next = append(next, currentRange)
			}
		}

		// This will load the next values to intRanges
		intRanges = intRanges[:0]
		intRanges, next = next, intRanges

		// If our map is empty no merges were performed so we have the final state and break the loop
		if len(mergedRanges) == 0 {
			break
		}

	}

	// Now for each range we only need to see how many elements are contained inside and add them to our counter
	for _, r := range intRanges {
		freshIDs += r[1] - r[0] + 1
	}
	fmt.Printf("Part 2 - Number of fresh IDs in ranges: %d\n", freshIDs)
}
