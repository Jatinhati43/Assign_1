package main

import (
	"fmt"
	"github.com/Jatinhati43/Assign_1/internal/number"
)

func main() {
	arr := []int{10, 5, 20, 8, 20,67,334,323,23,2,3,2,2,32565,45,7}

	secondMax, err := numbers.FindSecondHighest(arr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("The second highest number is: %d\n", secondMax)
}
