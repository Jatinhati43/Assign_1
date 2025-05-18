package numbers

import (
	"fmt"
	"math"
)


func FindSecondHighest(nums []int) (int, error) {
	if len(nums) < 2 {
		return 0, fmt.Errorf("array must contain at least two elements")
	}

	highest := math.MinInt
	secondHighest := math.MinInt

	for _, num := range nums {
		if num > highest {
			secondHighest = highest
			highest = num
		} else if num > secondHighest && num != highest {
			secondHighest = num
		}
	}

	if secondHighest == math.MinInt {
		return 0, fmt.Errorf("no distinct second highest element found")
	}

	return secondHighest, nil
}
