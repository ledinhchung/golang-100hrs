package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func binarySearch(items []int, key int) int {
	left, right, loop := 0, len(items)-1, 1
	// Implement logic
	for left <= right {
		fmt.Println("Loop time: ", loop)
		loop++

		mid := left + (right-left)/2

		if items[mid] == key {
			return mid
		} else if items[mid] > key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	var datas []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Error converting string to int %v\n", err)
			continue
		}

		datas = append(datas, num)
	}

	fmt.Println("Input datas from file", datas)

	key := 8
	find := binarySearch(datas, key)

	if find != -1 {
		fmt.Printf("Found key %d at %dth of datas\n", key, find)
	} else {
		fmt.Printf("Key %d not found in list of datas\n", key)
	}
}
