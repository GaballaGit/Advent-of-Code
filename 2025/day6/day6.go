package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartOne(in *os.File) {
	var myNums [][]int
	var operations [][]string

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		// trim := strings.TrimSpace(string(scanner.Bytes()))
		// fmt.Println(trim)
		// row := strings.Split(string(scanner.Bytes()), " ")
		row := strings.Fields(string(scanner.Bytes()))
		if row[0] == "+" || row[0] == "*" {
			for _, elm := range row {
				operations = append(operations, []string{elm})
			}
			break
		}
		var tmp []int
		for _, elm := range row {
			val, _ := strconv.Atoi(elm)
			tmp = append(tmp, val)
		}
		myNums = append(myNums, tmp)
	}

	var ansArray []int
	fmt.Println(myNums, "\n", operations)
	for k, elm := range operations {
		tmp := 0
		switch elm[0] {
		case "+":
			for i := 0; i < len(myNums); i++ {
				fmt.Println("+", k, i)
				tmp += myNums[i][k]
			}
		case "*":
			tmp = 1
			for i := 0; i < len(myNums); i++ {
				fmt.Println("*", k, i, len(myNums))
				tmp *= myNums[i][k]
			}
		}
		ansArray = append(ansArray, tmp)
	}

	var ans int
	for _, elm := range ansArray {
		ans += elm
	}

	fmt.Println("Part 1 >>", ans)
}

func main() {
	input, _ := os.Open("./input.txt")
	defer input.Close()

	PartOne(input)
	//PartTwo(input)
}
