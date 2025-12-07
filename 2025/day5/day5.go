package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func PartOne(in *os.File) {
	reg := regexp.MustCompile("([0-9]+)-([0-9]+)")
	var ids []int
	var ranges [][2]int

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		raw := scanner.Bytes()
		line := make([]byte, len(raw))
		copy(line, raw)
		n := string(line)
		if !reg.MatchString(n) {
			id, _ := strconv.Atoi(n)
			ids = append(ids, id)
		} else {
			numRange := reg.FindStringSubmatch(n)
			start, _ := strconv.Atoi(numRange[1])
			end, _ := strconv.Atoi(numRange[2])
			ranges = append(ranges, [2]int{start, end})
		}
	}
	ans := 0
	for _, elm := range ids {
		for _, r := range ranges {
			if elm >= r[0] && elm <= r[1] {
				ans++
				break
			}
		}
	}

	fmt.Println(ans)

}

func main() {
	n, _ := os.Open("input.txt")
	defer n.Close()

	// m := n
	PartOne(n)
	// PartTwo(m)
}
