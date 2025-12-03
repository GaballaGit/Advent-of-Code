package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func PartOne(input *os.File) {
	scanner := bufio.NewScanner(input)

	ans := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		ans += GetBattery(line)
	}
	fmt.Println(ans)
}

func GetBattery(line []byte) int {
	val := 0
	idx := -1
	startAt := 0
	sB := false

	for i := 1; i <= 2; i++ {
		bestIdx := -1
		curHigh := -1
		for j := startAt; j < len(line); j++ {
			if line[j] > byte(curHigh)+'0' && j != idx {
				curHigh = int(line[j]) - '0'
				bestIdx = j
			}
		}

		if sB {
			curHigh *= 10
		}

		val += curHigh
		idx = bestIdx
		startAt = bestIdx
		if i == 1 {
			if idx == len(line)-1 {
				startAt = 0
				sB = true
			} else {
				val *= 10
			}
		}
	}
	fmt.Println(val)
	return val
}

// ========================= PART TWO =========================

func PartTwo(input *os.File) {
	scanner := bufio.NewScanner(input)

	ans := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		ans += GetGigaVoltage(line)
		fmt.Println("=====================================")
	}
	fmt.Println(ans)
}

func GetGigaVoltage(line []byte) int {
	var val int = 0
	seen := make(map[int]bool)
	var idx []int
	startAt := 0

	for i := 1; i <= 12; i++ {
		bestIdx := -1
		curHigh := -1
		for j := startAt; j < len(line); j++ {
			if line[j] > byte(curHigh)+'0' && !seen[j] {
				curHigh = int(line[j]) - '0'
				bestIdx = j
			}
		}

		seen[bestIdx] = true
		idx = append(idx, bestIdx)
	}

	slices.Sort(idx)
	for j, elm := range idx {
		fmt.Println(val, elm, int(line[elm])-'0')
		var e int = int(line[elm]) - '0'
		val += e
		fmt.Println("e to val:", e, "->", val)
		if j != len(idx)-1 {
			val *= 10
		}
	}
	fmt.Println(idx)
	fmt.Println(val)
	return val
}

func main() {
	input, _ := os.Open("./input.txt")
	defer input.Close()

	// PartOne(input)
	PartTwo(input)
}
