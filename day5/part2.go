package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func isInRanges(num int, ranges []Range) bool {
	for _, r := range ranges {
		if num >= r.start && num <= r.end {
			return true
		}
	}
	return false
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
			// Empty line marks transition to phase 2
			break
		}

		// Parse range like "3-5"
		ends := strings.Split(line, "-")
		start, _ := strconv.Atoi(ends[0])
		end, _ := strconv.Atoi(ends[1])

		// check for overlap
		freshRanges = append(freshRanges, Range{start: start, end: end})

		// sort the slice of ranges
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("number of fresh ingredients:", len(list_fresh))
	fmt.Println("Done")
}
