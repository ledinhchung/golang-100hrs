package main

import (
	"fmt"
)

// Create a list for string from user input
func main() {
	var list [10]string
	var isExist string
	var value string
	index := 0

	for {
		fmt.Println("Please enter value: ")
		fmt.Scan(&value)

		if value != "" {
			list[index] = value
			index++
		}

		fmt.Println("Do you want to Exist insert? (Y/n)")
		fmt.Scan(&isExist)
		if isExist == "Y" || index > 10 {
			fmt.Printf("Your input: %v\n", list)
			break
		}
	}
}
