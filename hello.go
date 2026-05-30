package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "John"
	const ticketNumber = 34
	var remainingTickets = 100

	fmt.Printf("Remaining tickets: %d tickets available.\n", remainingTickets)
	fmt.Println("Hello, World!")
	fmt.Println("Welcome", name, "to Go programming.")
	fmt.Printf("Your  name on the ticket is: %s\n", name)

	//Ask user for the name and store it in the userName variable
	var userName string
	fmt.Print("Please Enter Your Name:-")
	fmt.Scanln(&userName)

	//Now print the userName variable
	fmt.Printf("Your name is: %s\n", userName)

	var userEmailId string
	fmt.Print("Please Enter Your Email Id:-")
	fmt.Scanln(&userEmailId)
	fmt.Printf("Your email id is: %s\n", userEmailId)
	//%T is used to print the type of the variable

	fmt.Printf("The type of userName variable is: %T\n", userName)
	fmt.Printf("The type of userEmailId variable is: %T\n", userEmailId)
	fmt.Printf("The type of ticketNumber variable is: %T\n", ticketNumber)
	fmt.Printf("The type of remainingTickets variable is: %T\n", remainingTickets)

	userAge := 25 //Using short variable declaration to declare and initialize userAge variable
	fmt.Printf("Your age is: %d\n", userAge)
	fmt.Printf("The type of userAge variable is: %T\n", userAge)
	//Printing the memory address of userAge variable using %p format specifier and & operator to get the address of the variable
	fmt.Printf("The memory address of userAge variable is: %p\n", &userAge)

	//Now we will update the remainingTickets variable by subtracting 1 from it and then print the updated value of remainingTickets variable

	remainingTickets = remainingTickets - 1
	fmt.Printf("Remaining tickets: %d tickets available.\n", remainingTickets)
	fmt.Printf("Thank you %s for booking a ticket. You will receive a confirmation email at %s.\n", userName, userEmailId)
	var bookings = []string{"Nana", "Praxy", "Jimmy"}

	fmt.Printf("The first booking is: %s\n", bookings[0])
	fmt.Println(bookings)
	fmt.Println(len(bookings))

	userName1 := "Alex"
	bookings = append(bookings, userName1) //append function is used to add a new element to the end of the slice

	fmt.Println(bookings) //The Slice is a data structure that can hold a variable number of elements. It is similar to an array but it can grow and shrink in size. The append function is used to add new elements to the end of the slice. In this case, we are adding the userName1 variable to the bookings slice. After appending, we print the updated bookings slice which now includes "Alex".

	//We will ask for the user's name and then we will check if the user's name is in the bookings slice. If it is, we will print a message saying "Welcome back, [user's name]!" Otherwise, we will print a message saying "Welcome, [user's name]!". in the loop
	for index, value := range bookings {
		fmt.Println(index, value) //Now the index is also being printed along with the value of each element in the bookings slice. The index starts from 0 and goes up to the length of the bookings slice minus 1. The value is the actual element in the bookings slice at that index.
		//The range keyword is used to iterate over the elements of a slice. It returns the index and the value of each element in the slice. In this case, we are using the blank identifier (_) to ignore the index and only get the value of each element in the bookings slice. We then print each value in the bookings slice.
		//Here we are asking the user for their name and then we are checking if the user's name is in the bookings slice. If it is, we will print a message saying "Welcome back, [user's name]!" Otherwise, we will print a message saying "Welcome, [user's name]!".
		//Will do later

	}
	if userName == "Nana" || userName == "Praxy" || userName == "Jimmy" || userName == "Alex" {
		fmt.Printf("Welcome back, %s!\n", userName)
		return //Stopping the execution of the program after printing the welcome back message
	} else {
		fmt.Println("Lol")

	}
	var checkerName string
	fmt.Print("Please Enter Your Name:-")
	fmt.Scanln(&checkerName)

	//Check the validation of the usser's name
	var lengthOfName = len(checkerName)
	if lengthOfName < 2 {
		fmt.Println("Name is too short. Please enter a valid name.")
		return //Stopping the execution of the program after printing the error message
	}
	var checkValidEmail string
	fmt.Print("Please Enter Your Email Id:-")
	fmt.Scanln(&checkValidEmail)
	if !strings.Contains(checkValidEmail, "@") {
		fmt.Println("Invalid email address. Please enter a valid email address.")
		return
	}
	//Checking the validity of the user's  email address by checking if it contains the "@" symbol. If it does not contain the "@" symbol, we will print an error message and stop the execution of the program. If it does contain the "@" symbol, we will print a message saying "Thank you for providing a valid email address." and continue with the rest of the program.

}
