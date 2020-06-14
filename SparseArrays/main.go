package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func matchingStrings(strings []string, queries []string) []int32 {
	counter := make(map[string]int)
	for _, str := range strings {
		if _, ok := counter[str]; ok {
			counter[str] += 1
		} else {
			counter[str] = 1
		}
	}

	counts := make([]int32, len(queries))
	for i, query := range queries {
		if count, ok := counter[query]; ok {
			counts[i] = int32(count)
		}
	}
	return counts
}

func main(){
	// Parse input
	strings, queries := parseInputLines()

	// Execute solution
	ans := matchingStrings(strings, queries)
	fmt.Println("output:", ans)
}

func parseInputLines() ([]string, []string) {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// read input ns: number of strings
	temp := readLine(reader)
	nLines, err := strconv.ParseInt(temp, 10, 64)
	if err != nil {
		panic("error reading input ns: number of strings")
	}

	var arrStr = make([]string, 0, nLines)
	for i := 0; i < int(nLines); i++ {
		item := strings.TrimSpace(readLine(reader))
		arrStr = append(arrStr, item)
	}

	// read input nq: number of queries
	temp = readLine(reader)
	nQuery, err := strconv.ParseInt(temp, 10, 64)
	if err != nil {
		panic("error reading input nq: number of queries")
	}

	var arrQ = make([]string, 0, nQuery)
	for i := 0; i < int(nQuery); i++ {
		item := strings.TrimSpace(readLine(reader))
		arrQ = append(arrQ, item)
	}

	return arrStr, arrQ
}

func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return "" }
	return string(bstr) 
}