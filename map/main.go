package main
import "fmt"
func main(){
	fmt.Println("i want to learn to the basic of the map:")
	var myMap = make(map[string]int)
	myMap["apple"] = 10
	myMap["banana"] = 20
	myMap["grapes"] = 30
	fmt.Println("map is : ",myMap)
	fmt.Println("length of map is : ",len(myMap))
	value, exists := myMap["banana"]
	if exists {
		fmt.Println("value for 'banana' is :", value)
	} else {
		fmt.Println("'banana' does not exist in the map")
	}
    myMap["banana"] = 25
	fmt.Println("updated map is : ",myMap)
	delete(myMap, "grapes")
	fmt.Println("map after deleting 'grapes':", myMap)

	for key , value := range myMap {
		fmt.Printf("key : %s , value : %d\n",key,value)
	}

	personAge := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Charlie": 35,
	}
	fmt.Println("-----------personAge map is :",personAge)

}