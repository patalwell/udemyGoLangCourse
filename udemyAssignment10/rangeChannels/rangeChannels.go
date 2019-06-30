package main

import "fmt"

//function to send an int to our channel
//since we are iterating close our channel after we are done
func sendInt(c chan <- int){
	for i:=1; i <= 5; i++ {
		c <- i
	}
	close(c)
}

//function to return an int from out channel ranging over our channel
func receiveInt(c <- chan int){
	for v:= range c{
		fmt.Println(v)
	}

}

func main() {

	//create a generic channel
	c := make(chan int)

	//send and int to our channel
	go sendInt(c)

	//receive an int from our channel
	//We dont need a waitGroup here, since this code will block until the channel has something to receive
	receiveInt(c)

}
