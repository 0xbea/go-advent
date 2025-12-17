package main

import (
	"bufio"
	"fmt"
	"math"
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
		idRanges := strings.Split(line, ",") // 11-22
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
			// todo: dont bother with range where both have odd digits and same number
			// invalid ID:  any ID which is made only of some sequence of digits repeated twice.
			// generate invalid IDs for a range
			for i := start; i <= end; i++ {
				// each range has a set
				// dont bother with IDs with odd number digits
				numDigits := len(strconv.Itoa(i))
				if numDigits%2 != 0 {
					continue
				}
				// if even number of digits, take first half and duplicate
				firstHalf := i / (10 * (numDigits - 1))
				checkId := firstHalf*int(math.Pow10(numDigits-1)) + firstHalf
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
