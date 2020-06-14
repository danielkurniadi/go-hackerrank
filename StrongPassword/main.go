package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
)

// Password Criteria Contains ALL:
// numbers = "0123456789"
// lower_case = "abcdefghijklmnopqrstuvwxyz"
// upper_case = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// special_characters = "!@#$%^&*()-+"
// Its length is at least 6

// minimumNumber check how many min character
// to add to make strong password
func minimumNumber(n int, password string) int32 {
	isnum, lower, upper, special := false, false, false, false
	for _, ch := range password {
		if isnum || checkIsNumbr(ch) { isnum = true }
		else if lower || checkIsLower(ch) { lower = true }
		else if upper || checkIsUpper(ch) { upper = true }
		else if special || checkSpecial(ch) { special = true }
	}
	var count int32
	for _, ok := range []bool{isnum, lower, upper, special} {
		if !ok { count ++ }
	}
	if count > 6 - n {
		return count
	}
	return 6 - n
}

func checkIsNumbr(ch rune) bool { return '0' <= ch && ch <= '9' }
func checkIsLower(ch rune) bool { return 'a' <= ch && ch <= 'z' }
func checkIsUpper(ch rune) bool { return 'A' <= ch && ch <= 'Z' }
func checkSpecial(ch rune) bool {
	switch {
	case 33 <= ch && ch <= 47:
		return true
	case 58 <= ch && ch <= 64:
		return true
	case 91 <= ch && ch <= 96:
		return true
	case 123 <= ch && ch <= 126:
		return true
	}
	return false
}

func main() {
	// Read input line
	input := parseInputLine()

	// Execute solution
	ans := minimumNumber(len(input), input)
	fmt.Println("outtput:", ans)
}

func parseInputLine() string {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	_ = readLine(reader) // ignore input n: length of password

	// read input string
	str := strings.TrimSpace(readLine(reader))
	return str
}

func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return  "" }
	return string(bstr)
}