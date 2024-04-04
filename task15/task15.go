package main

import (
	"fmt"
	"math/rand"
	"unsafe"
)

var justString string

func createHugeString(size int) string {
	letter := "abcdefghijklmnopqrstuvwxyz"
	runes := []rune(letter)
	buffer := make([]rune, size)
	for i := range buffer{
		buffer[i] = runes[rand.Intn(len(runes))]
	}

	return string(buffer)
}


func someFunc() {
	v := createHugeString(1 << 10)
	//use []rune for avoid problem with not ASCII symbol
  	justString = string(append([]rune{}, []rune(v)[:100]...))
	// copy
  	//justString = v[:100]
	fmt.Println(unsafe.StringData(v))
	

}

func main() {
	someFunc()
	fmt.Println(unsafe.StringData(justString))
}
