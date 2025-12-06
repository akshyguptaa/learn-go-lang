package main

import "fmt"

func main() {

	//This following assignment is called "type interface"
	//Only works inside functions
	helloWorld := "Hello, World!"
	fmt.Println(helloWorld)

	//Assign a different value to the variable
	helloWorld = "World, Hello!"
	fmt.Println(helloWorld)

	//Variable declarations:
	//var variable type
	//var variable type = value
	//var variable = value
	//var variable1, variable2 type = value1, value

	var name string
	var age int = 24
	var weight = 62.65
	var width, height int = 3, 4
	var flag bool = true

	var days []string = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	var user = map[string]string{
		"name":   "John Doe",
		"gender": "Male",
		"email":  "john.doe@email.com",
	}
	//variable declaration using var keyword can be at package level or inside function.
	//Variables declared at the function level must be used. Otherwise,the compiler is going to throw an error during compilation.
}
