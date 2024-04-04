package main

import (
	"fmt"
	"strings"
)
// reverse words
func reverseWords(s string) string{
	r := strings.Fields(s)
	left, right := 0, len(r) - 1
	for left < right{
		r[left], r[right] = r[right], r[left] // swap last and first word
		left ++
		right -- 
	}
	return strings.Join(r, " ") //concatenation into a string
}

func main(){

	s := "snow dog sun"
	fmt.Println(reverseWords(s))


}