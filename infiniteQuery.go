package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Insert your code here
	var input string

	fmt.Println("Please key in any number")
	_, _ = fmt.Scanln(&input)

	i, err := strconv.ParseInt(input, 10, 0)

	if err != nil {
		fmt.Println("Please only key in number", err)
	}

	if i%2 == 0 {
		fmt.Println(i, "is an even number")
		if i > 9 {
			fmt.Println("And it has more than 1 digit")
		} else {
			fmt.Println("And it has only 1 digit")
		}
	} else if i%2 == 1 {
		fmt.Println(i, "is an odd number")
		if i > 9 {
			fmt.Println("And it has more than 1 digit")
		} else {
			fmt.Println("And it has only 1 digit")
		}
	}
}
