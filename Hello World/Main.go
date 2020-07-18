package main

import (
	"fmt"
	"os"
)

// Type of comment
/* Another type of comment*/

func getSquare() int {
	return 5 * 5
}

func main() {
	// declaring name variable
	name := "Shreyas"
	fmt.Println("Hello my name is " + name)

	// declaring variales

	/*var cards string := "ace it or leave it"*/
	cards := "ace of spade"
	cards = "Heart cards"
	fmt.Println(cards)

	// the simple function to find the square

	square := getSquare()
	fmt.Println(square)
	// Arrays and slice
	// should be of same type
	c := []string{"apple", "orange"}

	fmt.Println(c)

	c = append(c, "blueberry")
	fmt.Println(c)

	for i, cs := range c {
		fmt.Println(i, " ", cs)
	}
	fact := 1
	for i := 1; i <= 5; i++ {
		fact *= i
	}

	fmt.Println("Factorial of 5 =", fact)

	os.Exit(143)

}
