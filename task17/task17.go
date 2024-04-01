package main

import "fmt"

func binsearch(arr []int, target int) int{
	left := 0
	right := len(arr) - 1
	for left <= right{
		m := (left + right) / 2
		switch{
		case target == arr[m]:
			return m
		case arr[m] < target:
			left = m + 1
		case arr[m] > target:
			right = m - 1
		}
	}
	return -1
}

func main(){
	arr := []int{-10, 1, 3, 12, 16, 19}
	fmt.Println(binsearch(arr, 12))
}