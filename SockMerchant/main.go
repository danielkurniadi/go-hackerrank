package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sockMerchant(n int32, arr []int32) int32 {
	sockID2Count := make(map[int32]int32)
	for i := 0; i < len(arr); i++ {
		if _, ok := sockID2Count[arr[i]]; ok {
			sockID2Count[arr[i]] += 1
		} else {
			sockID2Count[arr[i]] = 1
		}
	}
	
	var totalPair int32
	for _, count := range sockID2Count {
		totalPair += count / 2
	}
	return totalPair
}

func main() {
	// Read input
	n, arr := parseInputLine()

	// Execute solution
	ans := sockMerchant(n, arr)
	fmt.Println("pairs of socks to match: ", ans)
}

func parseInputLine() (int32, []int32) {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// read input n
	n, err := strconv.Atoi(readLine(reader))
	if err != nil {
		panic(fmt.Errorf("error parsing input n"))
	}

	// reader input arr
	arTemp := strings.Split(readLine(reader), " ")

	var arr = make([]int32, 0, len(arTemp))
	for i := 0; i < len(arTemp); i++ {
		item, err := strconv.Atoi(arTemp[i])
		if err != nil {
			panic(fmt.Errorf("error parsing input arr of socks"))
		}

		arr = append(arr, int32(item))
	}
	return int32(n), arr
}

func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return "" }
	return strings.TrimRight(string(bstr), "\r\n")
}