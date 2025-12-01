package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PartOne(file *os.File) int {
	scanner := bufio.NewScanner(file)

	ans := 0
	pointer := 50
	for scanner.Scan() {
		line := scanner.Bytes()
		strNum := string(line[1:])

		num, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}

		inc := 1
		if line[0] == 'L' {
			inc *= -1
		}

		fmt.Println(ans, " | ", string(line[0]), num, pointer, "v")
		for range num {
			if pointer == 0 {
				ans++
			}
			pointer += inc
			if pointer == 100 && inc == 1 {
				pointer = 0
			} else if pointer == -1 && inc == -1 {
				pointer = 99
			}
		}
		fmt.Println(pointer)
		fmt.Println(ans)

		if pointer == 0 {
			ans++
		}
	}

	return ans
}

func PartTwo(file *os.File) int {
	scanner := bufio.NewScanner(file)

	ans := 0
	pointer := 50
	for scanner.Scan() {
		fmt.Println("SCANNN")
		line := scanner.Bytes()
		strNum := string(line[1:])

		num, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}

		inc := 1
		if line[0] == 'L' {
			inc *= -1
		}

		// mfw I have to learn proper modulo math
		for i := 0; i < num; i++ {
			pointer = (pointer + inc + 100) % 100
			fmt.Println(pointer)
			if pointer == 0 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	input, err := os.Open("./input.txt")
	// input, err := os.Open("./testCase2.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// pt1 := PartOne(input)
	pt2 := PartTwo(input)
	// fmt.Println("> PART 1 | ", pt1)
	fmt.Println("> PART 2 | ", pt2)
}
