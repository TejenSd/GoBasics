package main

import "fmt"

func main() {
	//while loop using for
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// //Infinite loop using for
	// for {
	// 	fmt.Println("This is an infinite loop")

	// }

	for i = 1; i <= 3; i++ {

		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// for loop witth range
	for i, v := range []int{1, 2, 3} {
		fmt.Println("Index:", i, "Value:", v)
	}

	for i := range 11 {
		if condition := i % 2; condition == 0 {
			continue
			
		}
		fmt.Println(i)
	}

}
