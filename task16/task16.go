package main

import (
	"fmt"
)

func quiksort(arr []int) []int{
	if len(arr) == 1 || len(arr) == 0{
		return arr
	}
	pivot := arr[len(arr) / 2] // select the pivot  as the middle value
	var left, right, res []int 
	for _, val := range arr{
		if val < pivot{
			left = append(left, val)
		} else if val > pivot{
			right = append(right, val)
		}
	}
	return append(append(append(res, quiksort(left)...), pivot), quiksort(right)...) // create new slice = quiksort(left) + pivot + quiksort(right)
}

func main(){
	l := []int{4, 5, 1, 7, 16, 15, 0, 9}
	fmt.Println(quiksort(l))
}