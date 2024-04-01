package main

import "fmt"

func removeElem(arr []int, ind int) []int{
	return append(arr[:ind], arr[ind + 1:]...)
}

func main(){
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} 
	fmt.Println(removeElem(arr, 2))
}