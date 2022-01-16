package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct { //blueprint
	firstName   string
	lastName    string
	birthDate   string
	createdDate time.Time //this is a custom type defined by go, from the built in time package
}

// creating a method for the User struct that can be used for the user struct specifically
func (user User) /* this connects this to the User struct here */ outputUserDetails() { //originally had user of type user for input - better to get the pointer here to not create another copy
	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthDate) //have access to fields stored inside the user struct
}

var reader = bufio.NewReader(os.Stdin)

func NewUser(fName string, lName string, bDate string) *User { //struct type as return type - also saying User here will be a pointer as opposed to the record itself
	created := time.Now()

	user := User{
		firstName:   fName,
		lastName:    lName,
		birthDate:   bDate,
		createdDate: created,
	}

	return &user
	// returns the pointer for user
}

func main() {
	var newUser User

	firstName := getUserData("Please enter your first name then hit enter: ")
	lastName := getUserData("Please enter your last name then hit enter: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY) then hit enter: ")

	newUser = *NewUser(firstName, lastName, birthdate)
	// dereferencing the User pointer for NewUser

	// newUser = User{} This creat,es an empty instance
	// If you don't use all values above the not used one will be created with default of empty string

	// ... do something awesome with that gathered data

	newUser.outputUserDetails()
	//needs to be called differently here as it's now a method
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.Replace(userInput, "\n", "", -1)

	return cleanedInput

}
