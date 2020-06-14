package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
	"strconv"
)

func aVeryBigSum(ar[]int64) int64 {
	var bigsum int64
	for i := 0; i < len(ar); i++ {
		bigsum += ar[i]
	}
	return bigsum
}

func main() {
	// Read input test
	input := parseInputLine()

	// Execute solution
	ans := aVeryBigSum(input)
	fmt.Println(ans)
}

func parseInputLine() []int64 {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	strconv.ParseInt(readLine(reader), 10, 64)

	arTemp := strings.Split(readLine(reader), " ")
	
	var input = make([]int64, 0, len(arTemp))
	for i := 0; i < len(arTemp); i++ {
		item, err := strconv.ParseInt(arTemp[i], 10, 64)
		if err != nil { panic(fmt.Errorf("error reading input: %v", arTemp)) }
		input = append(input, item)
	}
	return input
}

// readLine reads a single line and return empty
// string if it reads an EOF
func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return "" }
	return string(bstr)
}