package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string{
	r := strings.Fields(s)
	left, right := 0, len(r) - 1
	for left < right{
		r[left], r[right] = r[right], r[left]
		left ++
		right -- 
	}
	return strings.Join(r, " ")
}

func main(){

	s := "snow dog sun"
	fmt.Println(reverseWords(s))


}