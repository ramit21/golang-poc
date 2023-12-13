package commands

import "fmt"

// First letter of the method/field to be made upper case if it is to be made public to outside packages
// Every function must specify its return type.
func ProcessName string() {
	fmt.Printf("Processing name....\n")
	return "processed"
}