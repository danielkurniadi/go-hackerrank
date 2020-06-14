package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

var precomputedMS = [][][]int32 {
	{ {8, 1, 6}, {3, 5, 7}, {4, 9, 2} },
	{ {6, 1, 8}, {7, 5, 3}, {2, 9, 4} },
	{ {4, 9, 2}, {3, 5, 7}, {8, 1, 6} },
	{ {2, 9, 4}, {7, 5, 3}, {6, 1, 8} },
	{ {8, 3, 4}, {1, 5, 9}, {6, 7, 2} },
	{ {4, 3, 8}, {9, 5, 1}, {2, 7, 6} },
	{ {6, 7, 2}, {1, 5, 9}, {8, 3, 4} },
	{ {2, 7, 6}, {9, 5, 1}, {4, 3, 8} },
}

func calculateAbsDiff(matA, matB [][]int32) int32 {
	var absdiff int32
	for i := 0; i < len(matA); i++ {
		for j := 0; j < len(matA[0]); j++ {
			absdiff += abs(matA[i][j] - matB[i][j])
		}
	}
	return absdiff
}

func formingMagicSquare(square [][]int32) int32 {
	var minCost int32 = 1 << 10
	for _, magicSquare := range precomputedMS {
		cost := calculateAbsDiff(square, magicSquare)
		minCost = MinInt(minCost, cost)
	}
	return minCost
}

func abs(x int32) int32 {
	if x < 0 { return -x }
	return x
}

func MinInt(x, y int32) int32 {
	if x < y { return x }
	return y
}

func main() {
	// Read input line
	square := parseInputLine()
	
	// Execute solution
	minCost := formingMagicSquare(square)
	fmt.Println("minimum cost to change to a magic square", minCost)
}

func parseInputLine() [][]int32 {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// assume 3x3 square inputs
	var square = make([][]int32, 0, 3)

	for i := 0; i < 3; i++ {
		row := make([]int32, 3)
		arTemp := strings.Split(readLine(reader), " ")
		
		for j := 0; j < len(arTemp); j++ {
			item, err := strconv.Atoi(arTemp[j])
			if err != nil {
				panic(fmt.Errorf(
					"error reading square element to int : %s", arTemp[j]))
			}
			
			row[j] = int32(item)
		}
		square = append(square, row)
	}
	return square
}

func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return "" }
	return strings.TrimRight(string(bstr), "\r\n")
}
