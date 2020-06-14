package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Day of Programmer is 256th day of the year
const day256 = 256

func dayOfProgrammer(year int32) string {
	day, calendar := day256, generateMonthCalendar(int(year))
	month := 0
	for ; month < len(calendar); month++ {
		if day >= calendar[month] { 
			day -= calendar[month]
		} else {
			break
		}
	}
	// check if year 1918 and this is February
	if year == 1918 && month == 1 {
		day += 13 // skip to 14th of Feb after 31st of Jan
	}
	return fmt.Sprintf("%02d.%02d.%04d", day, month+1, year)
}

func generateMonthCalendar(year int) []int {
	calendar := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year < 1918 && (year % 4 == 0) {
		// check gregorian calendar
		// modify February to 29 days on leap day
		calendar[1] = 29
	} else if year > 1918 && (year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)) {
		// check julian calendar
		// modify February to 29 days on leap day
		calendar[1] = 29
	} else if year == 1918 {
		// historical calendar change.
		calendar[1] = 15
	}
	return calendar
}

func main() {
	inputTests := parseInputLine()

	for _, year := range inputTests {
		ans := dayOfProgrammer(year)
		fmt.Println("day of programmer:", ans)
	}
}

func parseInputLine() []int32 {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	N, err := strconv.Atoi(readLine(reader))

	if err != nil { panic(fmt.Errorf("error parsing input n: number of test cases")) }

	// read years line by line
	years := make([]int32, 0, N)
	for i := 0; i < N; i++ {
		item := readLine(reader)
		if item == "" { break }

		year, err := strconv.Atoi(item)
		if err != nil { panic(fmt.Errorf("error parsing input year to int: %s", item)) }

		years = append(years, int32(year))
	}
	return years
}

func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return "" }
	return strings.TrimSpace(string(bstr))
}
