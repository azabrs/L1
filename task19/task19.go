package main

import "fmt"
// reverse word
func reverse(s string) string{
	r := []rune(s)
	left, right := 0, len(r) - 1
	for left < right{
		r[left], r[right] = r[right], r[left] // swap last and first letter
		left ++
		right -- 
	}
	return string(r)
}

func main(){
	s := "главрыба"
	fmt.Println(reverse(s))

}