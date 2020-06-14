package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func mergeRegion(region, query []int32) [][]int32 {
    leftInter, rightInter := MaxInt(region[0], query[0]), MinInt(region[1], query[1])
    leftOuter, rightOuter := MinInt(region[0], query[0]), MaxInt(region[1], query[1])

    var newRegions = make([][]int32, 0, 3)
	var leftOutVal, rightOutVal int32

    if leftOuter == region[0] {
        leftOutVal, rightOutVal = region[2], query[2]
    } else {
        leftOutVal, rightOutVal = query[2], region[2]
    }
    
    if leftOuter < leftInter {
        newRegions = append(newRegions, []int32{leftOuter, leftInter-1, leftOutVal})
    }

    newRegions = append(newRegions, []int32{leftInter, rightInter, region[2] + query[2]})

    if rightOuter > rightInter {
        newRegions = append(newRegions, []int32{rightInter+1, rightOuter, rightOutVal})
    }

	fmt.Println(region, query)
	fmt.Println(newRegions)
    return newRegions
}

func checkIntersect(region []int32, query []int32) bool {
    if (region[0] <= query[0] && query [0] <= region[1]) { return true }
    return (region[0] <= query[1] && query[1] <= region[1])
}

func arrayManipulation(n int32, queries[][]int32) int64 {
    regions := make([][]int32, 0, len(queries) * 10)
	regions = append(regions, queries[0])
	queries = queries[1:]

    for len(queries) > 0 {
		query := queries[0]
		queries = queries[1:]

        nRegion := len(regions)
		
		for i := 0; i < nRegion; i++ {
            region := regions[i]
			
			// check if there is any intersect
            if checkIntersect(region ,query) {
                newRegions := mergeRegion(region, query)
				
				regions[i] = newRegions[0]
				if len(newRegions) == 1 { break }
				
				// there are 3 new regions
				query = newRegions[1]

				if len(newRegions) >= 3 {
					queries = append(queries, newRegions[2])
				}
            }
		}
		regions = append(regions, query)
		fmt.Println("regions:", regions)
		fmt.Println("queries:", queries)
    }

	fmt.Println(regions)

	var maxVal int64
	return maxVal
}

func MaxInt(x, y int32) int32 {
    if x > y { return x }
    return y
}

func MinInt(x, y int32) int32 {
    if x < y { return x }
    return y
}

func main() {
	// Parse input
	n, queries := parseInputLines()

	// Execute solution
	ans := arrayManipulation(int32(n), queries)
	fmt.Println("Max sum of array manipulation:", ans)
}

func parseInputLines() (int, [][]int32) {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nm := strings.Split(readLine(reader), " ")
	n, err := strconv.Atoi(nm[0])
	m, err := strconv.Atoi(nm[1])
	if err != nil {
		panic("error reading input n, m")
	}

	var queries = make([][]int32, 0, m)
	for i := 0; i < m; i++ {
		items := strings.Split(readLine(reader), " ")

		start, _ := strconv.Atoi(items[0])
		end, _ := strconv.Atoi(items[1])
		add, _ := strconv.Atoi(items[2])

		query := []int32{int32(start), int32(end), int32(add)}
		queries = append(queries, query)
	}
	return n, queries
}

func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return "" }
	return strings.TrimRight(string(bstr), "\r\n")
}
