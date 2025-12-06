package main 

import "fmt"

func main() {

	//Simple if statement
	var flag bool = true;
	if flag == true {
		fmt.Println("Flag is true")
	}

	//Simple if else statement
	var age  int = 18
	if age >= 18 {
		fmt.Println("User is an adult")
	} else {
		fmt.Println("User is a minor")
	}

	//Simple else if statement
	var score int = 50
	if score >= 75 {
		fmt.Println("Distinction")
	} else if score >= 60 {
		fmt.Println("First Class")
	} else {
		fmt.Println("Pass")
	}
}