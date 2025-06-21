package main

import "fmt"

func main() {
	//while loop using for
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//Infinite loop using for
	for {
		fmt.Println("This is an infinite loop")

	}

}
