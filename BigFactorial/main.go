package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int) string {
	bignum := make([]int, 0, 100)
	bignum = append(bignum, 1)
	for k := 2; k <= n; k++ {
		temp := make([]int, len(bignum))

		for num, offset := k, 0; num > 0; num, offset = num/10, offset+1 {
			for i := 0; i < len(bignum); i++ {
				if i + offset == len(temp) - 1 { temp = append(temp, 0) }
				
				temp[i + offset] += (num % 10) * bignum[i]

				if temp[i+offset] < 10 { continue }
				
				temp[i+offset+1] += temp[i+offset] / 10
				temp[i+offset] %= 10
			}
		}
		bignum = temp
	}
	numstr := make([]rune, len(bignum))
	for i := 0; i < len(bignum); i++ {
		numstr[i] = rune(bignum[len(bignum)-i-1]) + '0'
	}
	return strings.TrimLeft(string(numstr), "0")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    n, err := strconv.Atoi(readLine(reader))
    checkError(err)

	ans := extraLongFactorials(n)
	fmt.Printf("Factorial %d! = %s\n", n, ans)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
