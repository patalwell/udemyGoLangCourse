package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

//Exercise 2:
// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("The toJSON function failed to marshall: %v", err)
	}
	return bs, nil
}


//Exercise 3
//Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a
// value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”.

type customErr struct {
	s string
}

func (ce customErr) Error() string{
	return ce.s
}

func foo(e error){
	fmt.Println(e)
}

//Exercise 4
//use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your
//lat "50.2289 N"
//long "99.4656 W"

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		sre := fmt.Errorf("Square root error. Cannot perform negative comupations:", f)
		return 0, sqrtError{"50.2289 N","99.4656 W",sre}
	}
	return 42, nil
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	//Exercise 1:
	// Instead of using the blank identifier, make sure the code is checking and handling the error.
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatal("Unable to marshall p1 of type person", err)
	}
	fmt.Println(string(bs))

	//Exercise 2
	//Create a custom error message using “fmt.Errorf”.
	bs2,err := toJSON(p1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs2))

	//Exercise 3
	//Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a
	// value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”.
	customError := customErr{"Custom Error"}
	foo(customError)

	//Exercise 4:
	answer, err := sqrt(1.0)
	if err != nil{
		log.Println(err)
	}
	fmt.Println(answer)

	//Exercise 5:

}


