package main

import (
	"bufio"
	"os"
)

func PartTwo(in *os.File) {
	scanner := bufio.NewScanner(in)
	var nums [][]string

	for scanner.Scan() {
		line := scanner.Bytes()

		newNum := ""
		for k := range line {

		}

	}
}

func main() {
	in, _ := os.Open("./input.txt")
	defer in.Close()

	PartTwo(in)

}
