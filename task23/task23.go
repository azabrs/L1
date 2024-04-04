package main

import "fmt"

func removeElem(arr []int, ind int) []int{
	return append(arr[:ind], arr[ind + 1:]...) // delete the element with index ind while preserving the order of the original array
}

func main(){
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} 
	fmt.Println(removeElem(arr, 2))
}