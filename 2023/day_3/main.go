package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func parser() solution {
	f, err := os.Open("./2023/day_3/input.txt")
	if err != nil {
		log.Fatal("Not file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var cur_num number
	nums := make([]number, 0, 100)
	sym := make(map[point]struct{}, 100)
	gears := make(map[point]struct{}, 100)

	for row := 0; scanner.Scan(); row++ {
		for col, ch := range scanner.Text() {
			if unicode.IsDigit(ch) {
				if cur_num.value != 0 {
					cur_num.addDigit(row, col, int(ch-'0'))
				} else {
					cur_num = newNum(row, col, int(ch-'0'))
				}
			} else {
				if cur_num.value != 0 {
					nums = append(nums, cur_num)
					cur_num = number{}
				}
				if ch != '.' {
					sym[point{row, col}] = struct{}{}
					if ch == '*' {
						gears[point{row, col}] = struct{}{}
					}
				}
			}
		}
	}
	return solution{nums, sym, gears}
}

func part1() int {
	var total int
	solution := parser()
outer:
	for _, num := range solution.nums {
		for k := range num.pos {
			if _, contains := solution.sym[k]; contains {
				total += num.value
				continue outer
			}
		}
	}
	return total
}

func part2() int {
	var total int
	solution := parser()

gears_loop:
	for gear := range solution.gears {
		matches := make([]int, 0, 1000)
		for _, num := range solution.nums {
			if _, contains := num.pos[gear]; contains {
				if len(matches) == 2 {
					continue gears_loop
				}
				matches = append(matches, num.value)
			}
		}
		if len(matches) == 2 {
			total += matches[0] * matches[1]
		}
	}
	return total
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

type solution struct {
	nums  []number
	sym   map[point]struct{}
	gears map[point]struct{}
}

type point struct {
	row, col int
}

type number struct {
	value int
	pos   map[point]struct{}
}

func (n *number) addDigit(row, col, val int) {
	n.value = n.value*10 + val
	n.pos[point{row - 1, col + 1}] = struct{}{}
	n.pos[point{row + 1, col + 1}] = struct{}{}
	n.pos[point{row, col + 1}] = struct{}{}
}

func newNum(row, col, val int) number {
	return number{
		value: val,
		pos: map[point]struct{}{
			{row - 1, col - 1}: {},
			{row - 1, col + 1}: {},
			{row - 1, col}:     {},
			{row + 1, col - 1}: {},
			{row + 1, col + 1}: {},
			{row + 1, col}:     {},
			{row, col - 1}:     {},
			{row, col + 1}:     {},
		},
	}
}
