package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./2023/day_2/input.txt")
	if err != nil {
		log.Fatal("Not file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var part1, part2 int
	for scanner.Scan() {
		line := scanner.Text()
		g := strings.Split(line, ": ")
		gameId := strings.Split(g[0], " ")[1]
		setsSlice := strings.Split(g[1], "; ")
		possible := true
		min_cubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,	
		}
		for _, setStr := range setsSlice {
			colorMap := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}
			cubes := strings.Split(setStr, ", ")
			for _, cube := range cubes {
				c := strings.Split(cube, " ")
				num, err := strconv.Atoi(c[0])
				if err != nil {
					log.Fatalf("wrong string num: %s", c[0])
				}
				colorMap[c[1]] += num
			}
			for k, v := range colorMap {
				min_cubes[k] = max(min_cubes[k], v)
			}
			if colorMap["red"] > 12 || colorMap["green"] > 13 || colorMap["blue"] > 14 {
				possible = false
			}
		}
		if possible {
			gNum, err := strconv.Atoi(gameId)
			if err != nil {
				log.Fatalf("Wrong gameId: %s", gameId)
			}
			part1 += gNum
		}
		part2 += min_cubes["red"] * min_cubes["green"] * min_cubes["blue"]
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
