package main

import (
	"fmt"
)


//function to send an int to our channel
func sendInt(c chan <- int, sendVal int){
	c <- sendVal
}

//function to return an int from out channel
func receiveInt(c <- chan int){
		fmt.Println(<-c)

}

func main() {

	//create a generic channel
	c := make(chan int)

	//send and int to our channel
	go sendInt(c, 43)

	//receive an int from our channel
	//We dont need a waitGroup here, since this code will block until the channel has something to receive
	receiveInt(c)

}
