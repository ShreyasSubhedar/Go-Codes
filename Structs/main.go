package main

import "fmt"

// struct to resprest the person
type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// // 1st way
	// alex := person{"Alex", "Anderson"}
	// // 2nd way
	// Jack := person{firstName: "Jack", lastName: "Maa"}
	// // 3rd way
	// var john person
	// // It'll assign the zero value here
	// /* string -> ""
	//    int -> 0
	//    float ->0
	//    bool ->false
	// */
	// john.firstName = "John"
	// john.lastName = "Carder"
	// fmt.Println(alex, Jack)
	// fmt.Println(john)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jimparty@gmail.com",
			zipcode: 431001,
		},
	}
	fmt.Println(jim)

}
