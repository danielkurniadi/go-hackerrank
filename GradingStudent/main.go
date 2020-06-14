package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

func gradingStudents(grades []int32) []int32 {
	for i := 0; i < len(grades); i++ {
		if grades[i] < 38 { continue }
		grades[i] = roundToNextFive(grades[i])
	}
	return grades
}

func roundToNextFive(grade int32) int32 {
	rem := grade % 5
	if rem > 2 {
		grade = grade - rem + 5
	}
	return grade
}

func main() {
	// Parse input line
	inputs := parseInputLine()

	// Execute solution
	ans := gradingStudents(inputs)
	fmt.Println(ans)
}

func parseInputLine() []int32 {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// read number of grades
	n, err := strconv.Atoi(readLine(reader))
	if err != nil {
		panic(fmt.Errorf("error pasing input n to int: number of grades"))
	}

	var grades = make([]int32, 0, n)
	for i := 0; i < n; i++ {
		item, err := strconv.Atoi(readLine(reader))
		if err != nil {
			panic(fmt.Errorf("error pasing input grade to int: %d", item))
		}

		grades = append(grades, int32(item))
	}

	return grades
}

func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return "" }
	return strings.TrimSpace(string(bstr))
}
