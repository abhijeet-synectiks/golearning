package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define the Person struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Create an instance of Person
	person := Person{
		Name: "Alice",
		Age:  30,
	}

	// Use Encoder to encode the struct to JSON and write to standard output (os.Stdout)
	fmt.Println("Encoding to JSON:")
	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(person)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}

	// JSON data to decode
	jsonData := `{"name":"Bob","age":25}`

	// Use Decoder to decode JSON data from a string
	fmt.Println("Decoding from JSON:")
	var personFromJSON Person
	decoder := json.NewDecoder([]byte(jsonData))
	err = decoder.Decode(&personFromJSON)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	// Print the decoded struct
	fmt.Printf("Decoded Struct: %+v\n", personFromJSON)
}
