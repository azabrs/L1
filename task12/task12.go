package main

import "fmt"

// create ser
func set(s []string) []string {
	var res []string
	uniq := make(map[string]bool)

	for _, val := range s {
		uniq[val] = true
	}// remove duplicate

	for key := range uniq {
		res = append(res, key)
	}

	return res
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	setArr := set(arr)

	fmt.Printf("Origin arr: %v\nSet: %v\n", arr, setArr)
}