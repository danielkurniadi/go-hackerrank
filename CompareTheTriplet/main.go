package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
	var alicePoint, bobPoint int32 = 0, 0
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			alicePoint += 1
		} else if b[i] > a[i] {
			bobPoint += 1
		}
	}
	return []int32{alicePoint, bobPoint}
}

func main () {
	// Read test input
	a, b := parseInputLine()

	// Execute solution
	ans := compareTriplets(a, b)
	fmt.Println(ans)
}

func parseInputLine() ([]int32, []int32) {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	bTemp := strings.Split(strings.TrimSpace(readline(reader)), " ")

	var a, b []int32

	if len(aTemp) != 3 { 
		panic(fmt.Errof("error input line for array a: found %d items, expected 3", len(aTemp)))
	}

	if len(bTemp) != 3 { 
		panic(fmt.Errof("error input line for array b: found %d items, expected 3", len(bTemp)))
	}

	for i := 0; i < 3; i++ {
		item, err := strconv.ParseInt(aTemp[i])
		if err != nil {
			panic(fmt.Errorf("error parsing non int item %v", aTemp[i]))
		}
		a = append(a, item)
	}

	for i := 0; i < 3; i++ {
		item, err := strconv.ParseInt(bTemp[i])
		if err != nil {
			panic(fmt.Errorf("error parsing non int item %v", bTemp[i]))
		}
		b = append(b, item)
	}

	return a, b
}
