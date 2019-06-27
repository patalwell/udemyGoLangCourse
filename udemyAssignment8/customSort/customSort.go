package main

import (
	"fmt"
	"sort"
)

//Person Object
type Person struct {
	First string
	Age   int
}

//Type we want to sort on, e.g. Age
type ByAge []Person

//Overide the sort.Interface for []Person by Age field
//Since these methods are a receiver to ByAge they are of type ByAge
func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i,j int) {
	a[i], a[j] = a[j], a[i]
}
//Sort by ASC
func (a ByAge) Less(i,j int) bool {
	return a[i].Age < a[j].Age
}

//Overide String()
func (a Person) String() string {
	return fmt.Sprintf("Name: %s Age: %d",a.First, a.Age)
}

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)

	//Sort by Age Asc
	sort.Sort(ByAge(people))

	fmt.Println(people)
}
