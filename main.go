package main

import "fmt"

func main() {
	fmt.Println("Advent of Code Day 1")

	entries := []int{1721, 979, 366, 299, 675, 1456}
	t := 0
	for _, s := range entries {
		t += s + s
		fmt.Println(s, "+", s, "=", t)
		if s == 2020 {
			fmt.Println(t, "+", s, "= 2020")
		}
	}

}
