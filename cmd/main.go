// Package main demonstrates usage of the numbers package to find the second highest number in a slice.
//
// This program creates an integer slice, calls FindSecondHighest from the numbers package,
// and prints the second highest distinct number or an error message.
// @author Jatin Hati
// @date 2025-05-18
package main

import (
	"fmt"
	"github.com/Jatinhati43/Assign_1/internal/numbers"
)

func main() {
	arr := []int{10, 5, 20, 8, 20, 67, 334, 323, 23, 2, 3, 2, 2, 32565, 45, 7}

	secondMax, err := numbers.FindSecondHighest(arr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("The second highest number is: %d\n", secondMax)
}
