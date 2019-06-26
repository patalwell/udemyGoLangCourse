/*

Pointers allow you to share a value stored in some memory location. Use pointers when
you don’t want to pass around a lot of data you want to change the data at a location
Everything in Go is pass by value. Drop any phrases and concepts you may have from other languages.
Pass by reference, pass by copy - forget those phrases. “Pass by value.” That is the only phrase you need to know
and remember. That is the only phrase you should use. Pass by value.
Everything in Go is pass by value. In Go, what you see is what you get (wysiwyg). Look at what is happening.
That is what you get.

*/


package main

import "fmt"

var p *int

var a int = 5

func main() {

	//This is set to nil
	fmt.Println(p)
	//Displays a pointer to an int
	fmt.Printf("%T\n", p)

	i := 42
	//Prints the value of i
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	//Prints the address of i in memory
	fmt.Println(&i)

	p = &i
	//Prints the address of i
	fmt.Println(p)
	//Prints the value of i
	fmt.Println(*p)

	//Setting the value of i through a pointer
	*p = 21
	fmt.Println(i)

	//Printing the value of a
	fmt.Println(a)
	//Printing the value of a pointer to a
	fmt.Println(&a)
	a = 34
	fmt.Println(a)

}
