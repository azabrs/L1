package main

import (
	"fmt"
	"time"
)

func sleep(n int){
	timeout := time.After(time.Second * time.Duration(n))
	<- timeout // wait time.Second * time.Duration(n) second
}

func main(){
	start := time.Now()
	sleep(2)
	fmt.Println("Program was asleep",  time.Since(start))
	
}