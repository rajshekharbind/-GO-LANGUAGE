package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("i want to learn consepts of string in golang:")
	var str string = "Hello, golang"
	fmt.Println("String value:", str)
	fmt.Printf("Type of str: %T\n", str)

	// String concatenation
	var str1 string = "Hello, "
	var str2 string = "World!"
	concatenatedStr := str1 + str2
	fmt.Println("Concatenated String:", concatenatedStr)

	// String length
	strLength := len(str)
	fmt.Println("Length of str:", strLength)
	// Accessing individual characters
	firstChar := str[0]
	fmt.Printf("First character of str: %c\n", firstChar)
	// Iterating over a string
	fmt.Println("Characters in str:")
	for i, char := range str {
		fmt.Printf("Index %d: %c\n", i, char)
	}


	data := "ram.sharma.dinesh.com.in.pk"
	parts := strings.Split(data,".")
	fmt.Println("Splitted parts:", parts)
	for i, part := range parts {
		fmt.Printf("Part %d: %s\n", i, part)
	}

	joinedStr := strings.Join(parts, "-")
	fmt.Println("Joined String with '-':", joinedStr)


	prefix := "Hello"
	suffix := "Golang"
	fmt.Println("Does str start with 'Hello'? :", strings.HasPrefix(str, prefix))
	fmt.Println("Does str end with 'Golang'? :", strings.HasSuffix(str, suffix))

	substr := "golang"
	fmt.Println("Does str contain 'golang'? :", strings.Contains(str, substr))
	index := strings.Index(str, "golang")
	fmt.Println("Index of 'golang' in str:", index)

	upperStr := strings.ToUpper(str)
	lowerStr := strings.ToLower(str)
	fmt.Println("Uppercase String:", upperStr)
	fmt.Println("Lowercase String:", lowerStr)

	trimmedStr := strings.TrimSpace("   Hello, Golang!   ")
	fmt.Println("Trimmed String:", trimmedStr)

	replacedStr := strings.ReplaceAll(str, "golang", "Go")
	fmt.Println("Replaced String:", replacedStr)

	splitN := strings.SplitN(data, ".", 3)
	fmt.Println("SplitN parts (n=3):", splitN)

	
}