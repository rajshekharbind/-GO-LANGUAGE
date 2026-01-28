package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	fmt.Println("hi name is : ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("name is %s\n",name )


	reader := bufio.NewReader(os.Stdin)
	name2, _ := reader.ReadString('\n')
	fmt.Printf("Hello Mr %s", name2)
}
