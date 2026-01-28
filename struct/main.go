package main
import "fmt"

type Person struct {
		name string
		age  int
		city string
	}
type Contact struct {
	   email string
	   phone string

}
type Address struct {
	   street string
	   city   string
	   zip    string
}
type Employee struct {
	id int
	name  string
	contact int
	address string
}

func main(){
	fmt.Println("i want to learn to the basic of the struct concept:")
	var raj Person 
    fmt.Println("raj person : ",raj) 
	raj.name = "Raj"
	raj.age = 25
	raj.city = "Mumbai"
	fmt.Println("after assigning values to raj struct :",raj)
    //2nd method of declaring and initializing struct
	Person1 := Person{"Anita", 30, "Delhi"}
	fmt.Println("Person1 struct is :",Person1)
	//3rd method of declaring and initializing struct
	Person2 := Person{age: 28, name: "Vikram", city: "Bangalore"}
	fmt.Println("Person2 struct is :",Person2)
	//accessing struct fields
	fmt.Println("Name of Person1 is :",Person1.name)
	fmt.Println("Age of Person2 is :",Person2.age)
	//modifying struct fields
	Person2.city = "Chennai"
	fmt.Println("after modifying city of Person2 :",Person2)
	//new keyword to create struct
	Person3 := new(Person)
	Person3.name = "Sita"
	Person3.age = 22
	Person3.city = "Kolkata"
	fmt.Println("Person3 struct created using new keyword :",*Person3)




	var employee1 Employee
	employee1.id = 123
	employee1.name= "raj"
	employee1.contact = 9087654321
	employee1.address = "bhagalpur"
	fmt.Println("employee1 struct is :",employee1)


	var contact1 Contact
	contact1.email = "raj@example.com"
	contact1.phone = "9087654321"
	fmt.Println("contact1 struct is :",contact1)


	var address1 Address
	address1.street = "MG Road"
	address1.city = "Bangalore"
	address1.zip = "560001"
	fmt.Println("address1 struct is :",address1)

	
}