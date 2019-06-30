package main

import "fmt"

//send ints to our channels
func send(e, o, q chan<- int){
	for i:=0; i < 25; i++ {
		//if the int is even send to even chan
		if i%2 == 0 {
			e <- i
			//if the int is odd send to odd chan
		} else {
			o <- i
		}
	}
		q <- 0
}

//receive our ints from our channels
func receive(e, o, q <-chan int) {
	for {
		select {
			case v := <-e :
				fmt.Println("Value from even channel: ", v)
			case v := <- o:
				fmt.Println("Value from odd channel: ", v)
			case v := <- q:
				fmt.Println("Value from quit channel: ", v)
			return
		}
	}
}


func main() {
	//even int channel
	evenCh := make(chan int)
	oddChan := make(chan int)
	quitChan := make(chan int)

	//send
	go send(evenCh, oddChan, quitChan)

	//receive
	receive(evenCh, oddChan, quitChan)

	//Warning
	fmt.Println("About to exit...")
	
}
