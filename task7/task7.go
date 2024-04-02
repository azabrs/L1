package main

import (
	"fmt"
	"sync"
)


type Data_map struct{
	data map[int]int
	sync.RWMutex
}

func New_data() Data_map{
	return Data_map{
		data: make(map[int]int),
	}
}

func(d *Data_map) compet_read(key int) int{
	d.RLock()
	buf := d.data[key]
	d.RUnlock()
	return buf
}

func(d *Data_map)compet_write(key, val int){
	d.Lock()
	d.data[key] = val
	d.Unlock()
}

func main(){
	data := New_data()
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){

		defer wg.Done()
		for i := 0; i < 10; i++{
			fmt.Println(i, " - WRITE")
			data.compet_write(i, i * i)

		}
	}()

	wg.Add(1)
	go func(){

		defer wg.Done()
		for i := 0; i < 10; i++{
			buf := data.compet_read(i)
			fmt.Println("FIRST READ - ", buf)
		}
	}()


	wg.Add(1)
	go func(){

		defer wg.Done()
		for i := 0; i < 10; i++{
			buf := data.compet_read(i)
			fmt.Println("SECOND READ - ", buf)
		}
	}()
	fmt.Println("check")
	wg.Wait()
}