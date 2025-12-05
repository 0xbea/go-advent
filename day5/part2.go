package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func compareRanges(a, b Range) int {
	if a.start < b.start {
		return -1
	} else if a.start > b.start {
		return 1
	} else {
		return 0
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var freshRanges []Range

	scanner := bufio.NewScanner(file)

	// Phase 1: Read ranges of fresh ingredients
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Empty line marks end
			break
		}

		// Parse range like "3-5"
		ends := strings.Split(line, "-")
		start, _ := strconv.Atoi(ends[0])
		end, _ := strconv.Atoi(ends[1])

		// check for overlap
		overlapDetected := false
		for i, r := range freshRanges {
			// Four cases: contains, contained, overlaps left, overlaps right
			if start <= r.end && end >= r.start {
				if start < r.start {
					freshRanges[i].start = start
				}
				if end > r.end {
					freshRanges[i].end = end
				}
				overlapDetected = true
				break
			}
		}
		if !overlapDetected {
			freshRanges = append(freshRanges, Range{start: start, end: end})
			fmt.Printf("Added new range: %d-%d\n", start, end)
		} else {
			fmt.Printf("Merged into existing range. Now: %v\n", freshRanges)
		}
		// sort the slice of ranges
		slices.SortFunc(freshRanges, compareRanges)

		// maybe merge overlapping ranges
		for i := 0; i < len(freshRanges)-1; {
			current := freshRanges[i]
			next := freshRanges[i+1]
			// handle range going into next range
			if current.end >= next.start {
				// extend current range
				freshRanges[i].end = freshRanges[i+1].end
				// remove next range
				newRanges := freshRanges[0 : i+1]
				newRanges = append(newRanges, freshRanges[i+2:]...)
				freshRanges = newRanges
				// don't increment i, check again in case there's another overlap
			} else {
				i++
			}
		}
	}

	// sum up ranges now
	total := 0
	for _, r := range freshRanges {
		total += (r.end - r.start + 1)
	}
	fmt.Println("Total fresh ingredients:", total)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	// fmt.Println(freshRanges)
	fmt.Println("Done")
}
