package main

import (
	"fmt"
	"math/big"
)

// add two numbers a + b
func add(a, b *big.Int) *big.Int{

	return new(big.Int).Add(a, b)
}

// sub two numbers a + b
func sub(a, b *big.Int) *big.Int{
	return new(big.Int).Sub(a, b)
}

// multiply two numbers a * b
func mul(a, b *big.Int) *big.Int{
	return new(big.Int).Mul(a, b)
}

// dvide two numbers a / b
func div(a, b *big.Int) *big.Int{
	return new(big.Int).Div(a, b)
}


func main(){
	a := big.NewInt(int64(1 << 24))
	b := big.NewInt(int64(1 << 22))
	fmt.Println("a -", a)
	fmt.Println("b -", b)
	fmt.Println("add -", add(a, b))
	fmt.Println("sub -", sub(a, b))
	fmt.Println("div -", div(a, b))
	fmt.Println("mul -", mul(a, b))
}