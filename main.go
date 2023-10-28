// ## Initialize ## \\

package main

import (
	"fmt"
)

// ## Declare variables here ## \\

var doingOperation bool = false
// make the contact directory
var directory map[string]any = make(map[string]any)

// ## Start main program ## \\

func createContact(name string, number string) bool {
	if directory[name] != nil {
		fmt.Println("This contact already exists!")
		return false
	}

	directory[name] = number

	fmt.Printf("Assigned %s to number %s\n", name, number)

	return true
}

func getContact(name string) any {
	return directory[name]
}

func editContact(name string, newNumber string) bool {
	if directory[name] == nil {
		fmt.Println("The contact you wanted to edit doesn't exist!")
		return false
	}

	directory[name] = newNumber
	fmt.Printf("Changed %s's number to %s\n", name, newNumber)

	return true
}

func listContacts() {
	var output string = ""
	var cursor int = 0

	for Key, Value := range directory {
		cursor++

		output = output + fmt.Sprintf("Entry #%d - Name: %s, Number: %s \n", cursor, Key, Value)
	}

	if output == "" {
		fmt.Println("Listing 0 entries:\nNo records were put in yet!")
		return
	}
	fmt.Printf("Listing %d entries:\n", cursor)
	fmt.Println(output)
	return
}

func main() {
	createContact("Mark", "0217589187")
	createContact("Ethan", "054919495817")
	createContact("Alexandra", "01759174578")

	editContact("Ethan", "079174876")

	listContacts()
}
