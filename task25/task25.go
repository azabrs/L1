package main

import (
	"fmt"
	"time"
)

func sleep(n int){
	t := time.After(time.Second * time.Duration(n))
	<- t
}

func main(){
	start := time.Now()
	sleep(2)
	fmt.Println("Program was asleep",  time.Since(start))
	
}