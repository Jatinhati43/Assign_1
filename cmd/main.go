package main

import (
	"fmt"
	"github.com/yourusername/Assignment_1/internal/numbers"
)

func main() {
	arr := []int{10, 5, 20, 8, 20}

	secondMax, err := numbers.FindSecondHighest(arr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("The second highest number is: %d\n", secondMax)
}
