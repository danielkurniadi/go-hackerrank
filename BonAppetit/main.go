package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

func bonAppetit(bills []int32, k, bon int32) {
	var shareBill int32
	for i := 0; i < len(bills); i++ {
		if i == int(k) { continue }
		shareBill += bills[i]
	}
	annaBill := shareBill / 2
	if annaBill != bon {
		fmt.Println(bon - annaBill)
	}
	fmt.Println("Bon Appetit")
}

func main() {
	// Read input line
	bills, k, bon := parseInputLine()

	// Execute solution
	bonAppetit(bills, k, bon)
}

func parseInputLine() ([]int32, int32, int32) {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nkStr := strings.Split(readLine(reader), " ")
	
	n, err := strconv.Atoi(nkStr[0])
	k, err := strconv.Atoi(nkStr[1])
	if err != nil {
		panic(fmt.Errorf("error parsing input n, m: number of bills and anna: %v", err))
	}

	var bills = make([]int32, 0, n)

	// Read input bills
	arTemp := strings.Split(readLine(reader), " ")
	for i := 0; i < len(arTemp); i++ {
		item, err := strconv.Atoi(arTemp[i])
		if err != nil {
			panic(fmt.Errorf("error parsing input bills: %v", err))
		}

		bills = append(bills, int32(item))
	}

	// read bon appetit
	bon, _ := strconv.Atoi(readLine(reader))

	return bills, int32(k), int32(bon)
}

func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return "" }
	return strings.TrimSpace(string(bstr))
}
