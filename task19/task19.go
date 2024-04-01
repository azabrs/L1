package main

import "fmt"

func reverse(s string) string{
	r := []rune(s)
	left, right := 0, len(r) - 1
	for left < right{
		r[left], r[right] = r[right], r[left]
		left ++
		right -- 
	}
	return string(r)
}

func main(){
	s := "главрыба"
	fmt.Println(reverse(s))

}