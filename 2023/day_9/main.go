package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	nums := make([][]int, 0, len(lines))
	for _, line := range lines {
		numsStrSlice := strings.Split(line, " ")
		n := make([]int, 0, len(numsStrSlice))
		for _, numStr := range numsStrSlice {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatalf("Wrong number: %s", numStr)
			}
			n = append(n, num)
		}
		nums = append(nums, n)
	}
	return nums
}

func allZero(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func process_line(numSlice []int) int {
	var count int
	for !allZero(numSlice) {
		l := len(numSlice)
		diffs := make([]int, 0, l)
		count += numSlice[l-1]

		for i := 1; i < l; i++ {
			diffs = append(diffs, numSlice[i]-numSlice[i-1])
		}
		numSlice = append(make([]int, 0, len(diffs)), diffs...)
	}
	return count
}

func part1(input string) int {
	numsSlice := parse(input)
	var ans int
	for _, numSlice := range numsSlice {
		ans += process_line(numSlice)
	}
	return ans
}

func part2(input string) int {
	numsSlice := parse(input)
	var ans int
	for _, numSlice := range numsSlice {
		slices.Reverse(numSlice)
		ans += process_line(numSlice)
	}
	return ans
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Not file")
	}
	input := string(data)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
