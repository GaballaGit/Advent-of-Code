package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type cord [2]int

func PartOne(grid [][]byte) {
	var startCord cord
	for i, elm := range grid[0] {
		if elm == byte('S') {
			startCord = [2]int{0, i}
		}
	}

	fmt.Println(startCord)
	var queue []cord
	split := 0
	queue = append(queue, startCord)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur[0] == len(grid)-1 {
			fmt.Println("BREAK")
			break
		}
		look := grid[cur[0]+1][cur[1]]

		// fmt.Println(cur, queue, string(look))
		switch look {
		case '^':
			if cur[1] > -1 && grid[cur[0]+1][cur[1]-1] != '|' {
				grid[cur[0]+1][cur[1]-1] = '|'
				queue = append(queue, [2]int{cur[0] + 1, cur[1] - 1})
			}

			if cur[1] < len(grid) && grid[cur[0]+1][cur[1]+1] != '|' {
				grid[cur[0]+1][cur[1]+1] = '|'
				queue = append(queue, [2]int{cur[0] + 1, cur[1] + 1})
			}

			split++
		case '.':
			grid[cur[0]+1][cur[1]] = '|'
			queue = append(queue, [2]int{cur[0] + 1, cur[1]})
		}
	}

	OutputMatrix(grid)
	fmt.Println(split)
}

func PartTwo(grid [][]byte) {
	var startCord cord
	for i, elm := range grid[0] {
		if elm == byte('S') {
			startCord = [2]int{0, i}
		}
	}

	// var dfs func(cord)
	// dfs = func(i cord) {
	// 	if i[0]+1 == len(grid) {
	// 		ans++
	// 		return
	// 	}

	// 	if grid[i[0]+1][i[1]] == '^' {
	// 		dfs([2]int{i[0] + 1, i[1] - 1})
	// 		dfs([2]int{i[0] + 1, i[1] + 1})
	// 	} else {
	// 		dfs([2]int{i[0] + 1, i[1]})
	// 	}
	// }

	// dfs(startCord)

	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[startCord[0]][startCord[1]] = 1

	for r, elm := range grid {
		if r+1 == len(grid) {
			break
		}
		for c, elm2 := range elm {
			if c+1 == len(grid[0]) {
				break
			}

			if elm2 == '^' {
				dp[r+1][c-1] += dp[r][c]
				dp[r+1][c+1] += dp[r][c]
			} else {
				dp[r+1][c] += dp[r][c]
			}
		}

	}

	ans := 0
	for _, elm := range dp[len(grid)-1] {
		ans += elm
	}

	fmt.Println(ans)
}
func OutputMatrix(grid [][]byte) {

	fn := "out.txt"
	file, _ := os.Create(fn)
	defer file.Close()

	for _, elm := range grid {
		for _, elm2 := range elm {
			_, _ = fmt.Fprintf(file, string(elm2))
		}
		_, _ = fmt.Fprintf(file, "\n")
	}
}

func main() {
	i, _ := os.Open("./input.txt")
	defer i.Close()

	scanner := bufio.NewScanner(i)
	var grid [][]byte
	for scanner.Scan() {
		raw := scanner.Bytes()
		line := make([]byte, len(raw))
		copy(line, raw)
		grid = append(grid, line)
	}

	start := time.Now()
	PartOne(grid)
	end1 := time.Since(start)
	start = time.Now()
	PartTwo(grid)
	end2 := time.Since(start)
	fmt.Println("Part 1 time:", end1)
	fmt.Println("Part 2 time:", end2)
}
