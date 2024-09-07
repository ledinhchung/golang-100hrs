package main

import "fmt"

/* take input from terminal then check if input string is palindrome or not
*
* Example:
* not -> false
* sos -> true
* pop -> true
 */

func main() {
	// Allow user input many time
	var input string
	for {
		fmt.Print("Please input string or \"exit\" to exit: ")
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		r := isPalinDrome(input)
		out := "not"
		if r {
			out = "the"
		}

		fmt.Printf("%s is %s palindrom string\n", input, out)
	}
}

func isPalinDrome(s string) bool {
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		if s[i] != s[mid-1-i] {
			return false
		}
	}

	return true
}
