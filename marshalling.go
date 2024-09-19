package main

import (
	"encoding/json"
	"fmt"
)

// Define the Person struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Create an instance of the Person struct
	person := Person{
		Name: "Alice",
		Age:  30,
	}

	// Marshalling: Convert the struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	// Print the JSON data as a string
	fmt.Println("Marshalled JSON:", string(jsonData))

	// Unmarshalling: Convert JSON back to a struct
	var personFromJSON Person
	err = json.Unmarshal(jsonData, &personFromJSON)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	// Print the struct obtained from unmarshalling
	fmt.Println("Unmarshalled Struct:", personFromJSON)
}
