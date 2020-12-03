package main

import "fmt"

func main() {
	fmt.Println("Advent of Code Day 1")

	var entries []int = []int{1721, 979, 366, 299, 675, 1456, 4, 4}

	for i, transaction1 := range entries {
		for j := i + 1; j < len(entries); j++ {
			transaction2 := entries[j]
			if transaction1+transaction2 == 2020 {
				fmt.Println(transaction1, "+", transaction2, "= 2020")
				fixreport := transaction1 * transaction2
				fmt.Println("Give the elves this number", fixreport)
			}
		}
	}

	/* This code runs correctly, lets see if we can make it cleaner.
	for i, transaction1 := range entries {
		for j, transaction2 := range entries {
			if transaction1+transaction2 == 2020 && j > i {
				fmt.Println(transaction1, "+", transaction2, "= 2020")
			}
		}
	}*/

}
