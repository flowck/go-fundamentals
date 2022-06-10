package main

import (
	"fmt"
)

func pointers(shouldRun bool) {
	if !shouldRun {
		fmt.Println("[POINTERS] - Skipping pointers topic")
		return
	}

	// Operations
	// 0 - Reserve space in memory
	var firstName *string = new(string)

	// Must print the address in memory holding the string "John"
	fmt.Println(firstName)

	// 1 - Dereference a pointer and assign a value to it
	*firstName = "John"

	// Prints the actual string stored in the address the pointer points to
	fmt.Println(*firstName)

	// 2 - Obtaining the memory address of an existing variable
	lastName := "Doe"
	lastNamePointer := &lastName

	fmt.Println(lastNamePointer, lastName)
	fmt.Println(*lastNamePointer, &lastName)
}

func main() {
	fmt.Println("Go fundamentals")

	pointers(true)
}
