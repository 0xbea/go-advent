package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Circuit struct {
	Points      []int
	TotalLength float64
}

// { Points: [0,1], TotalLength: 100 }

func calcDistance(x1, y1, z1, x2, y2, z2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2) + math.Pow(z1-z2, 2))
}

func main() {
	file, err := os.Open("test_input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// set initial/global variables
	lineNumber := 0
	pointsList := make(map[int][3]float64)       // 0: [x, y, z]
	distList := make([]Circuit, len(pointsList)) // [{ Points: [0,1], TotalLength: 100 }, { Points: [0,2], TotalLength: 200 }, ...]

	scanner := bufio.NewScanner(file)
	// process each line
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(coords[0], 64)
		y, _ := strconv.ParseFloat(coords[1], 64)
		z, _ := strconv.ParseFloat(coords[2], 64)
		pointsList[lineNumber] = [3]float64{x, y, z}
		lineNumber++
	}

	// calculate distances between all points
	for i := 0; i < len(pointsList)-1; i++ {
		for j := i + 1; j < len(pointsList); j++ {
			dist := calcDistance(pointsList[i][0], pointsList[i][1], pointsList[i][2],
				pointsList[j][0], pointsList[j][1], pointsList[j][2])
			p2pCircuit := Circuit{
				Points:      []int{i, j},
				TotalLength: dist,
			}
			distList = append(distList, p2pCircuit)
		}
	}

	// sort distList
	sort.Slice(distList, func(i, j int) bool {
		return distList[i].TotalLength < distList[j].TotalLength
	})

	// connect first limit pairs
	limit := 1000
	for i := 1; i < limit; i++ {
		for j := 0; j < i; j++ {
			// check previous circuits for points in current pair
			for _, pt := range distList[i].Points {
				for _, 
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Done")
}
