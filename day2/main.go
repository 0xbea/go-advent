package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("test_input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// set initial / global variables
	var invalidList []int // list of invalid IDs

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// process each line ... there is only one
		idRanges := strings.Split(line, ",")
		for _, idRange := range idRanges {
			startEnd := strings.Split(idRange, "-")
			start, err := strconv.Atoi(startEnd[0])
			if err != nil {
				print(err)
			}
			end, err := strconv.Atoi(startEnd[1])
			if err != nil {
				print(err)
			}
			fmt.Println(start, end)
			// invalid ID:  any ID which is made only of some sequence of digits repeated twice.
			// dont bother with IDs with odd number digits
			numDigits := start/10 + 1
			if numDigits%2 != 0 {
				if end/10+1 == numDigits {
					continue
				}
			}
			// otherwise generate invalid IDs for a range
			// split it in half and duplicate digits
			tenPower := numDigits - 1
			firstHalf := start / (10 * tenPower) // floor division
			// lastHalf := end / (10 * tenPower)
			for i := firstHalf; i <= end; i++ {
				checkId := i*10*tenPower + i
				fmt.Println("checking:", checkId)
				fmt.Println("numDigits", numDigits)
				if checkId > end {
					break
				}
				invalidList = append(invalidList, checkId)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println(invalidList)
	fmt.Println("Done")
}
