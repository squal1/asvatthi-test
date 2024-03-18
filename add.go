package main

import (
	"fmt"
)

	
func plus(a int, b int) int {
    return a + b
}

// Define a custom type called "Person"
type Person struct {
	Name    string
	Age     int
	Country string
}

// Method associated with the Person type to introduce the person
func (p Person) Introduce() {
	fmt.Printf("Hi, my name is %s. I'm %d years old and I'm from %s.\n", p.Name, p.Age, p.Country)
}

func main() {
	// Creating an instance of the Person type
	person1 := Person{
		Name:    "John",
		Age:     30,
		Country: "USA",
	}

	// Calling the Introduce method on person1
	person1.Introduce()

	// Creating another instance of the Person type
	person2 := Person{
		Name:    "Emily",
		Age:     25,
		Country: "Canada",
	}

	// Calling the Introduce method on person2
	person2.Introduce()
}
