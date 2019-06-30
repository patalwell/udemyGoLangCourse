package main

import "fmt"

//set ints to our channels
func send(e, o chan<- int, q chan<- bool){
	for i := 0; i < 10; i++ {
		//if the int is even send to even chan
		if i%2 == 0 {
			e <- i
			//if the int is odd send to odd chan
		} else {
			o <- i
		}
	}
	close(q)
}

//get ints from our channels
func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
			case v := <-e :
				fmt.Println("Value from even channel: ", v)
			case v := <- o:
				fmt.Println("Value from odd channel: ", v)
			case i, ok := <- q:
				//we can use , ok idiom to see if a channel is closed
				if !ok {
					fmt.Println("From comma ok", i, ok)
					//if we no longer have a value exit the look
					//e.g. if ok returns false
					return
				} else {
					fmt.Println("From comma ok", i)
				}
		}
	}
}


func main() {
	//even int channel
	evenCh := make(chan int)
	oddChan := make(chan int)
	quitChan := make(chan bool)

	//generic channel with , ok idiom
	c := make(chan int)

	//send
	go send(evenCh, oddChan, quitChan)

	//receive
	receive(evenCh, oddChan, quitChan)


	//comma, ok
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println("Here is the value of our channel:",v, "and here is the state:",ok)

	//Our channel is closed once we get an item
	v, ok = <-c
	fmt.Println("Here is the value of our channel:",v, "and here is the state:",ok)

	//Warning
	fmt.Println("About to exit...")

}