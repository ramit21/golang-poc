package main

import (
	"flag"
	"fmt"
	//"./commands"
)

//Including another package of the project

// Go provides a `flag` package supporting basic command-line flag parsing.

func main() {
	//Calling a method from a different file in the same package does not require an import
    //welcome()

    //Calling an exported method from a different file in a different package
    returnValFromFunc := ""
    //returnValFromFunc := commands.ProcessName()
    returnValFromFunc = "value updated without a := this time."
    fmt.Println(returnValFromFunc)

    //calling a receiver function on a reference of a custom type
    cards := newDeck()
    cards.print()


    // Basic flag declarations (passed from command line at time of execution) are available for string, integer, and boolean options. 
    // Here we declare a string flag `name` with a default value `"world"` and a short description.
    // This `flag.String` function returns a string pointer (not a string value)
    namePtr := flag.String("name", "World", "Default name (optional)")
    flag.Parse()
    fmt.Println("Hello", *namePtr)
}
