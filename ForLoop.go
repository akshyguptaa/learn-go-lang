package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var names = []string{"Tom", "John", "Harry"}

	fmt.Println("Simple for loop : ")
	//Simple for loop
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println("For loop with break: ")
	//For loop with break
	var id string = "Test123"
	var usernames []string = []string{"Test321", "JohnD", "Test123"}

	for i := 0; i < len(usernames); i++ {
		if usernames[i] == id {
			fmt.Println("Username found in  the list")
			fmt.Println(i)
			break
		}
	}

	fmt.Println("For loop with continue: ")
	//For loop with continue
	var evenCount int = 0
	var counts []int = []int{1, 2, 4, 7, 3, 567, 987, 234, 567, 213, 675987, 234, 213, 546, 76}
	for i := 0; i < len(counts); i++ {
		if counts[i]%2 != 0 {
			continue
		}
		evenCount++
	}
	fmt.Println("Total count of even numbers : ")
	fmt.Println(evenCount)

	fmt.Println("For loop as while loop: ")
	//For as while loop (Sort of)
	var limit int = 10
	var i int = 0
	for i <= limit {
		if i%2 != 0 {
			fmt.Println(i)
		}
		i++
	}

	fmt.Println("For loop with auto indexing (Range) : ")
	for _, item := range usernames {
		fmt.Println(item)
	}

	fmt.Println("For loop with auto indexing (Range) but for maps : ")
	var scores = map[string]int{
		"John":  60,
		"James": 66,
		"Jane":  78,
	}
	for name, score := range scores {
		fmt.Println(name, score)
	}

}
