// ## Initialize ## \\

package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"	
	"strings"
)

// ## Declare variables here ## \\

var doingOperation bool = false
var exiting bool = false
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

func readInput() string{
	reader := bufio.NewReader(os.Stdin)

	in, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	in = strings.TrimSuffix(in, "\n")
	in = strings.TrimSuffix(in, "\r")

	return in
}

func promptInput(){
	if doingOperation == true {
		return
	}

	fmt.Println("Type in what you want to do: [1:Create][2:Edit][3:List][4:Exit]")
	var readIn string = readInput()
	//fmt.Println(readIn)
	switch(readIn){
	case "1":
		fmt.Println("Type in the name:")
		var name string = readInput()
		fmt.Println("Type in the number:")
		var number string = readInput()

		createContact(name, number)
	case "2":
		fmt.Println("Type in the name:")
		var name string = readInput()
		fmt.Println("Type in the new number:")
		var number string = readInput()

		editContact(name, number)
	case "3":
		listContacts()
	case "4":
		fmt.Println("Exiting process now")
		exiting = true
	default:
		fmt.Println("Invalid operation")
	}	
}

func main() {
	for exiting == false {
		promptInput()
	}
}
