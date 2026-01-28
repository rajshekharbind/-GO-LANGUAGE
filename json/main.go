package main
import (
	"fmt"
    "encoding/json"

	
)
// ---------------- JSON MARSHALING ----------------
	type Person struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Email  string `json:"email"`
		Active bool   `json:"active"`
	}

func main() {
	fmt.Println("I want to learn concepts of json in Golang")
// ---------------- JSON MARSHALING ----------------
	person := Person{
		Name:   "John Doe",
		Age:    30,
		Email:  "john.doe@example.com",
		Active: true,
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println("JSON Marshaled Data:", string(jsonData))
	
// ---------------- JSON UNMARSHALING ----------------
	var person2 Person
	err = json.Unmarshal(jsonData, &person2)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Printf("Unmarshaled Struct: %+v\n", person2)


	// ---------------- PRETTY PRINTING JSON ----------------
	prettyJSON, err := json.MarshalIndent(person, "", "    ")
	if err != nil {
		fmt.Println("Error pretty printing JSON:", err)
		return
	}
	fmt.Println("Pretty Printed JSON:\n", string(prettyJSON))

	// ---------------- WORKING WITH JSON ARRAYS ----------------
	people := []Person{
		{Name: "Alice", Age: 25, Email: "alice@example.com", Active: true},
		{Name: "Bob", Age: 28, Email: "bob@example.com", Active: false},

	}
	peopleJSON, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error marshaling JSON array:", err)
		return
	}
	fmt.Println("JSON Array Marshaled Data:", string(peopleJSON))

	var people2 []Person
	err = json.Unmarshal(peopleJSON, &people2)
	if err != nil {
		fmt.Println("Error unmarshaling JSON array:", err)
		return
	}
	fmt.Printf("Unmarshaled JSON Array: %+v\n", people2)
		fmt.Println("Error creating HTTP request:", err)
		
}
