package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup


func main() {
	wg.Add(2)

	//directional buffered channel
	//e.g can only add items to the channel
	c := make(chan int, 2)
	//When reading channel or pointer type read left to righ
	// e.g. channel that receives values
	cr := make(chan <- int, 2)

	//e.g. channel that sends values
	cs := make(<- chan int, 2)

	go func(){
		fmt.Println("Routine receiving values on channel")
		cr <- 42
		cr <- 43
		c <- 44
		wg.Done()
	}()

	//This wont work!
	//fmt.Println(<- c)

	//This wont work unless we use a generic channel
	go func(){
		fmt.Println("Routine sending values to STDOUT")
		fmt.Println(<-c)
		wg.Done()
	}()


	//Get the channel type
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	//General to specific channels
	//e.g int to float is ok, float to int you'll lose data
	//cr = c
	//cs = c

	//Wait till routine runs prior to closing main()
	wg.Wait()
	
}
