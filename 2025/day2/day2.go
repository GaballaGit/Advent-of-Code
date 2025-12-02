package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOne(input []byte) {
	ranges := strings.Split(string(input), ",")

	// fmt.Println(ranges)
	reg, _ := regexp.Compile(`([0-9]+)-([0-9]+)`)

	ans := 0
	for _, elm := range ranges {
		subStr := reg.FindStringSubmatch(elm)
		ans += helper1(subStr[1:])
	}

	fmt.Println(ans)
}

func helper1(in []string) int {
	idx1, err := strconv.Atoi(in[0])
	if err != nil {
		panic(err)
	}
	idx2, err := strconv.Atoi(in[1])
	if err != nil {
		panic(err)
	}

	i := idx1
	n := idx2

	val := 0
	for i <= n {
		stringify := strconv.Itoa(i)

		if len(stringify)%2 == 0 {
			val += check1(stringify)
		}

		i++
	}

	return val
}

func check1(x string) int {
	mid := len(x) / 2

	for i := 0; mid < len(x); mid, i = mid+1, i+1 {
		if x[mid] != x[i] {
			return 0
		}
	}

	num, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	return num
}

// ========================== PART TWO ==========================

func PartTwo(input []byte) {
	ranges := strings.Split(string(input), ",")

	// fmt.Println(ranges)
	reg, _ := regexp.Compile(`([0-9]+)-([0-9]+)`)

	ans := 0
	for _, elm := range ranges {
		subStr := reg.FindStringSubmatch(elm)
		ans += helper2(subStr[1:])
	}

	fmt.Println(ans)
}

func helper2(in []string) int {
	idx1, err := strconv.Atoi(in[0])
	if err != nil {
		panic(err)
	}
	idx2, err := strconv.Atoi(in[1])
	if err != nil {
		panic(err)
	}

	i := idx1
	n := idx2

	val := 0
	for i <= n {
		stringify := strconv.Itoa(i)

		val += check2(stringify)
		// fmt.Println("=========================================")
		i++
	}

	return val
}

func check2(x string) int {
	n := len(x) / 2
	for i := 0; i < n; i++ {
		v := strings.Repeat(x[:i+1], len(x)/(i+1))
		// fmt.Println(v, "||", x)
		if v == x {
			num, err := strconv.Atoi(x)
			if err != nil {
				panic(err)
			}
			return num
		}
	}
	return 0
}

func main() {
	input, err := os.Open("./input.txt")
	// input, err := os.Open("./testCase.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	in, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}

	PartOne(in)
	PartTwo(in)
}
