package main

import (
	"fmt"
	"sync"
)

type inc_type struct{
	sync.RWMutex
	inc int
}

func New() inc_type{
	return inc_type{
		inc : 0,
	}
}

func (i *inc_type)Inc(){
	i.Lock()
	i.inc += 1
	i.Unlock()
}

func main(){
	c := New()
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.inc) 
}