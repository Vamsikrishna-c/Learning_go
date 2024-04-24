package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Declare a variable to store the user input
	var input string

	// Prompt the user for input
	fmt.Println("Enter something: ")

	// Read the user input using fmt.Scanln
	fmt.Scanln(&input)

	// Print the user input
	fmt.Println("You entered:", input)

	// //to read the whole line we use bufio package
	// textReader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter something: ")
	// //ReadString returns two values String and error
	// text, _ := textReader.ReadString('\n')
	// fmt.Println("You entered:", text)

}
