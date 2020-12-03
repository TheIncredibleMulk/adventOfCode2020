package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader, _ := os.Open("Trasnactions.txt")
	scanner := bufio.NewScanner(reader)
	var nums []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	seen := make(map[int]bool)

	for _, num := range nums {
		complement := 2020 - num
		_, ok := seen[complement]
		if ok {
			fmt.Println(complement * num)
			break
		}
		seen[num] = true
	}
}

/*
func main() {
	transactions, err := ioutil.ReadFile("Transactions.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s", content)
}
*/
