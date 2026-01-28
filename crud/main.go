package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ---------------- MODEL ----------------
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// ---------------- CREATE ----------------
func createTodo() {
	todo := Todo{
		UserID:    1,
		Title:     "Learn CRUD in Golang",
		Completed: false,
	}

	jsonData, _ := json.Marshal(todo)

	res, err := http.Post(
		"https://jsonplaceholder.typicode.com/todos",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("POST Error:", err)
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("CREATE Response:")
	fmt.Println(string(body))
}

// ---------------- READ ----------------
func readTodo() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("GET Error:", err)
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("\nREAD Response:")
	fmt.Println(string(body))
}

// ---------------- UPDATE ----------------
func updateTodo() {
	todo := Todo{
		UserID:    1,
		ID:        1,
		Title:     "Updated Todo Title",
		Completed: true,
	}

	jsonData, _ := json.Marshal(todo)

	req, _ := http.NewRequest(
		http.MethodPut,
		"https://jsonplaceholder.typicode.com/todos/1",
		bytes.NewBuffer(jsonData),
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("PUT Error:", err)
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("\nUPDATE Response:")
	fmt.Println(string(body))
}

// ---------------- DELETE ----------------
func deleteTodo() {
	req, _ := http.NewRequest(
		http.MethodDelete,
		"https://jsonplaceholder.typicode.com/todos/1",
		nil,
	)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("DELETE Error:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("\nDELETE Response Status:", res.Status)
}

// ---------------- MAIN ----------------
func main() {
	fmt.Println("🚀 CRUD Operations with JSONPlaceholder")

	createTodo()
	readTodo()
	updateTodo()
	deleteTodo()
}
