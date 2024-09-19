package main

import (
    "encoding/json"
    "fmt"
)

// Define a struct
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    // Create an instance of the struct
    p := Person{Name: "John Doe", Age: 30}

    // Marshal the struct to JSON
    jsonData, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the JSON data
    fmt.Println(string(jsonData))
}
