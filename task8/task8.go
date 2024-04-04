package main

import (
	"fmt"
	"strconv"
	)

func main(){
	fmt.Println("input bit position and value")
	var pos, val int
	fmt.Scan(&pos, &val)
	
	var n int64 = 0x12345678
	fmt.Printf("Input value %X", strconv.FormatInt(n, 16))
	mask := (1 << (pos - 1))
	if val == 1{
		n = n | int64(mask) // set pos bit to 1
	} else if val == 0{
		maskInv := ^int64((mask))
		n = n & maskInv // set pos bit to 0
	}
	fmt.Printf("Value after transformation %X", strconv.FormatInt(n, 16))
}
