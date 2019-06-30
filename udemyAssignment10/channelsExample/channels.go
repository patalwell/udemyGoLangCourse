package main

/*
	Note: Channels block, so you'll need a routine that unlocks the channel
	e.g. simply having chn <- 42 will not work without being wrapped in a routine since
	fmt.Println() also wants to access the channel
 */

import "fmt"

func main() {
	//use make() to create a channel of type chan <primitive type>
	chn := make(chan int)

	//use make() and param for num of int accepted to create a buffer channel which won't need a routine
	chnBuf := make(chan int, 1)

	//put an int on our chan
	// nel with a routine
	go func(){
		chn <- 42
	}()

	chnBuf <-43

	//this wont work
	//chnBuf <-44

	//Show channel value
	fmt.Println(<- chn)

	//show channel buf value
	fmt.Println(<- chnBuf)
	
}
