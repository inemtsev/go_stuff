package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 9}
	sNum := 7

	pos, iterations, ok := findNumber(nums, sNum)

	if ok {
		fmt.Printf("The number is in position %v. It took %v iterations", pos, iterations)
	} else {
		fmt.Println("Number wasnt found in the list")
	}

}

func findNumber(nums []int, sNum int) (mid int, iterations int, isFound bool) {
	minNum := 0
	maxNum := len(nums) - 1
	iterations = 0

	for {
		iterations++

		mid = (minNum + maxNum) / 2
		fmt.Println(mid)

		midNum := nums[mid]
		if midNum == sNum {
			isFound = true
			return
		}

		if minNum == maxNum {
			return
		}

		if sNum <= midNum {
			maxNum = mid
		} else {
			minNum = mid + 1
		}
	}

}
