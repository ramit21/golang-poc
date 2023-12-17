package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	//"./commands"
)

//Including another package of the project

// Go provides a `flag` package supporting basic command-line flag parsing.

func main() {

    // Basic flag declarations (passed from command line at time of execution) are available for string, integer, and boolean options. 
    // Here we declare a string flag `name` with a default value "world" and a short description.
    // This `flag.String` function returns a string pointer (not a string value), that's why we refer to the value using a *, ie *namePtr.
    namePtr := flag.String("name", "World", "Default name (optional)")
    flag.Parse()
    fmt.Println("Hello", *namePtr)

	//Calling a method from a different file in the same package does not require an import
    //welcome()

    //Calling an exported method from a different file in a different package
    returnValFromFunc := ""
    //returnValFromFunc := commands.ProcessName()
    returnValFromFunc = "value updated without a := this time."
    fmt.Println(returnValFromFunc)

    //calling a receiver function on a reference of a custom type
    cards := newDeck()
    cards.saveToFile("./deckfile.txt")
    cardFromfile := cards.newDeckFromFile("./deckfile.txt");
    cardFromfile.print()

    //passing struct using pointers for pass by reference
    jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{ //nested struct
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("jimmy")
    //Pointer shortcut: Alternatively, Go provides a cleaner way to write above 2 sentences in one line, 
    //just say below instead, and Go internally will treat it as the 2 lines above:
    //jim.updateName("jimmy")
	jim.print()

    //maps
    colorCodes := map[int]string { //this tells that keys are of type int and values are of type string
        1: "red",
        2: "blue",
    }
    colorCodes[3] = "Pink"
    //fmt.Println(colorCodes)
    //iterate over map keys:
    for colorCode, color := range colorCodes {
        fmt.Println("Color code for" , color , "is", colorCode)
    }
  
    //Making http calls
    resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(resp)
    fmt.Println("Now printing response body")
    io.Copy(os.Stdout, resp.Body)
    /* Above line is equivalent to:
    bs := make([]byte, 99999)
    resp.Body().Read(bs)
    fmt.Println(string(bs))
    */

    //Calling site checker go routines
    links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) //Making main calling code to wait for all threads to send value to this channel.
	}
}
