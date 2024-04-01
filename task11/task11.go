package main

import "fmt"

func intersection(set1, set2 []int) []int{
	intersec := make(map[int] bool)
	for _, val := range(set1){
		intersec[val] = true
	}
	for _, val := range(set2){
		intersec[val] = true
	}
	keys := make([]int, 0, len(intersec))
	for k := range intersec {
		keys = append(keys, k)
	}
	return keys

}


func main() {
	set1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	set2 := []int{4, 5, 6, 8, 1, 10, 11, 0}

	fmt.Println(intersection(set1, set2))
}