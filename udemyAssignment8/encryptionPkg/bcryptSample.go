package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	//Sample Password
	pass := "password123"
	wrongPass := "Changeme"

	//bcrypt's Generate Hash from Password
	//accepts []byte slice and Constant to determine the hashing strength
	//returns []byte slice and error
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Something went wrong during encryption: ", err)
	}

	//Here is the hashed string
	fmt.Println(string(hashedPass))

	err2 := bcrypt.CompareHashAndPassword(hashedPass,[]byte(wrongPass))

	if err2 == nil {
		fmt.Println("Passwords Match!")
	} else {
		fmt.Println("Wrong Password!")
	}

	
}
