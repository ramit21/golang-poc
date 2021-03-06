package main

import "fmt"
//Including another package of the project
import "./commands"

// Go provides a `flag` package supporting basic command-line flag parsing. 
import "flag"

func main() {
	//Calling a method from a different file in the same package
    welcome()

    //Calling an exported method from a different file in a different package
    commands.ProcessName()

    // Basic flag declarations are available for string, integer, and boolean options. 
    // Here we declare a string flag `name` with a default value `"world"` and a short description.
    // This `flag.String` function returns a string pointer (not a string value)
    namePtr := flag.String("name", "World", "Default name (optional)")
    flag.Parse()
    fmt.Println("Hello", *namePtr)
}
