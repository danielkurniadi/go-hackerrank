package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

func birthdayCakeCandles(candles []int32) int32 {
	tallest, countTallest := candles[0], 0
	for i := 0; i < len(candles); i++ {
		if tallest < candles[i] {
			tallest, countTallest = candles[i], 1
		} else if tallest == candles[i] {
			countTallest += 1
		}
	}
	return int32(countTallest)
}

func main() {
	// Read Parse line
	input := parseInputLine()

	// Execute solution
	ans := birthdayCakeCandles(input)
	fmt.Println("Candle (tallest) that can be blown:", ans)
}

func parseInputLine() []int32 {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// ignore input n: number of candles
	_ = readLine(reader)

	// read input candles height
	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	
	var arr = make([]int32, 0, len(arrTemp))
	for i := 0; i < len(arrTemp); i++ {
		item, err := strconv.ParseInt(arrTemp[i], 10, 64)
		if err != nil { panic(fmt.Errorf("error parsing input line %v", arrTemp)) }

		arr = append(arr, int32(item))
	}
	return arr
}

func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return "" }
	return string(bstr) 
}