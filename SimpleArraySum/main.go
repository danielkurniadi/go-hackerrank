package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strconv"
	"strings"
)

func main() {
	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)
	defer writer.Flush()

	// Parse input
	nums := parseInputLine()

	// Execute solution
	ans := simpleArraySum(nums)

	fmt.Fprintf(writer, "%v\n", ans)
}

func simpleArraySum(nums []int32) int32 {
	var sum int32
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func parseInputLine() []int32 {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// read number of items in test
	strconv.ParseInt(readLine(reader), 10, 64)

	// read test input
	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < len(arTemp); i++ {
		item, err := strconv.ParseInt(arTemp[i], 10, 64)
		if err != nil { 
			panic(fmt.Errorf("error parsing input item %d\n", item))
		}
		ar = append(ar, int32(item))
	}
	return ar
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}