# Pointers

A pointer holds the address of a variable in memory.
One can say that it points to the address in memory rather than an actual value;

## Examples

```go
package main

import "fmt"

func main() {

	// Creating a pointer
	var fullNamePtr *string = new(string)
	fmt.Println(fullNamePtr) // Prints the address
	
	// Storing value in the pointer's address by dereferencing it
	*fullNamePtr = "John Doe"
	fmt.Println(fullNamePtr, *fullNamePtr)
	
	// Getting the pointer of an existing variable
	country := "Brazil"
	fmt.Println(&country)
}
```