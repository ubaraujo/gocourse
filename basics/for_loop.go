package basics

import "fmt"

func main() {

	// Simple iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Iterate over collection
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	// Using break with loops
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // continue the loop but skip the rest of lines/statements
		}
		fmt.Println("Odd Number: ", i)

		if i == 5 {
			break
		}
	}

	// Outer Loop
	rows := 5
	for i := 1; i <= rows; i++ {
		//inner loop for spaces before stars
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		// inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // move to the next line

	}

	// iterating over integers directly
	for i := range 10 {
		fmt.Println(10 - i)
		i++
	}
	fmt.Println("We have a lift off!")

}
