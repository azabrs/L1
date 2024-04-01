package main

import (
	"fmt"
	"time"
)

var start time.Time



func receiver(ch chan string){
	for{
		fmt.Println("send message", time.Since(start))
		time.Sleep(time.Millisecond * 500)
		ch<-"message"


	}

}

func main(){
	ch := make(chan string)
	var N int
	fmt.Scan(&N)
	timer := time.After(time.Second * time.Duration(N))
	start = time.Now()
	go receiver(ch)
	for {
		select{
			case <-timer:
				fmt.Println("program finished", time.Since(start))
				close(ch)
				return
			case val := <-ch:
				fmt.Println(val, "was received - ", time.Since(start))
		}
	}

}
