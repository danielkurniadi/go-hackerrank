package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var directions = [][]int{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	{-1, -1}, {-1, 1}, {1, 1}, {1, -1},
}

func queensAttack(n int, k int, rq int, cq int, obstacles [][]int) int {
	var obstacleSet = make(map[[2]int]bool)
	for _, obstacle := range obstacles {
		posy, posx := obstacle[0], obstacle[1]
		obstacleSet[[2]int{posy, posx}] = true
	}	
	totalCount := 0
	for _, dir := range directions {
		dy, dx := dir[0], dir[1]
		totalCount += countAttackedInDirection(dy, dx, rq, cq, n, obstacleSet)
	}
	return totalCount
}

func countAttackedInDirection(dy, dx, rq, cq, n int, obstacleSet map[[2]int]bool) int {
	i, j := rq+dy, cq+dx
	count := 0
	for 1 <= i && i <= n && 1 <= j && j <= n {
		if _, ok := obstacleSet[[2]int{i, j}]; ok { break }
		i, j = i+dy, j+dx
		count++
	}
	return count
}

func main() {
	// Read parse input line
	n, k, rq, cq, obstacles := parseInputLine()

	// Execute solution
	ans := queensAttack(n, k, rq, cq, obstacles)
	fmt.Println("queen attack cells", ans)
}

func parseInputLine() (n, k, rq, cq int, obstacles [][]int) {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	
	var err error
	posTmp := strings.Split(readLine(reader), " ")
	n, err = strconv.Atoi(posTmp[0])
	k, err = strconv.Atoi(posTmp[1])

	if err != nil {
		panic(fmt.Errorf("error parsing input n, k"))
	}

	posTmp = strings.Split(readLine(reader), " ")
	rq, err = strconv.Atoi(posTmp[0])
	cq, err = strconv.Atoi(posTmp[1])

	if err != nil {
		panic(fmt.Errorf("error parsing input r_q, c_q"))
	}

	for i := 0; i < k; i++ {
		posTmp := strings.Split(readLine(reader), " ")
		row, err := strconv.Atoi(posTmp[0])
		col, err := strconv.Atoi(posTmp[1])
		if err != nil {
			panic(fmt.Errorf("error parsing input obstacles row, col "))
		}

		obstacles = append(obstacles, []int{row, col})
	}
	return
}

func readLine(reader *bufio.Reader) string {
	bstr, _, err := reader.ReadLine()
	if err == io.EOF { return "" }
	return strings.TrimRight(string(bstr), "\r\n")
}
