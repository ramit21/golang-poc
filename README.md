# golang-poc

Installing Golang: https://golang.org/doc/install

Learn about Golang: https://golang.org/

Steps to run the poc: (on Windows)

1) Ensure that Golang is installed and environment variables are set as explained in the installation link above

2) Git clone this repo

3) Execute the command 'go build' to create the executable

4) Run the executable: 'demo.exe'. Name of the executable is taken as specified in the go.mod file

5) Run the executable passing runtime arguments: 'demo -name=Your_name'

6) Run 'golang-poc -h' to see the help menu with all the run time arguments defined for the application

GO 1.5 onwards supports cross compilation (https://dave.cheney.net/2015/03/03/cross-compilation-just-got-a-whole-lot-better-in-go-1-5)
To run the poc on OSX platform, follow the above steps with following exceptions:
In Step 3, generate the shell script using the command: 'env GOOS=darwin GOARCH=386 go build'

Now execute the script as './golang-poc -name=Your_name'
  

## Theory  

### Go Cli
1. go build: compiles a bunch of go source code files.
2. go run <file-name>: compiles and executes a given go file. If using functions from other go files, give them all in run command: eg. 'go run main.go deck.go'
3. go fmt: format all files in current directory.
4. go install: install a package.
5. go get: download raw code of someone else's package.
6. go test: run any tests associated with the current project.

### package

Package serves as the workspace within a project. ie. grouping of code with similar purpose together. 
All go files that belong to same package should have same 
'package <package-name>' at the top.

Types of packages:
1. Executable: eg package main. Generates a file on 'go build' command, that we can run. Inside main package, you also need a func called main that acts as entrypoint of the executable.
2. Reusable: Code used as helpers, ie. reusable code. These can further be classified as part of reusable package, and std library.
Standard library packages can be checked at golang.org/pkg

### Variables

Variables in Golang are statically typed, like in Java, and not Dynamic like javascript.

However, you can leave it for Golang compiler to figure out the data type:

```
var msg1 string = "first msg" //you specify the variable type
msg2 := "second message" //compiler resolves this declaration to type string
msg2 = "third message" //note that := can be used only once at time of declaration
```

Variables can be declared outside functions, but can't be given a value there. Value should be assigned inside a function only.

### Pass by value vs pass by reference and pointers

Everything in Golang is pass by value. Even if you pass a pointer, and print its address using &, it will be different from original value. Just that both pointers would be pointing to the same underlying object, so it would behave as pass by reference.

```
func main(){
    name:= "abc"
    namePointer:=&name
    fmt.Println(namePointer) //prints different value than below
    printPointer(namePointer)
}

func printPointer(namePointer *string) {
    fmt.Println(&namePointer)
}

```

**Pass by reference**(those who behave as if pass by reference): slices, maps, channels, pointers, functions.

**Pass by value**: Basic types: int, float, string, bool, and structs. To make these objects pass by reference, Golang provides pointers.

```
&varName -> gives memory address of the value this variable is pointing at
*pointer -> Gives the value this memory address is pointing to
```

### Array vs Slices

Array is of fixed length, whereas slices are arrays that can grow or shrink.
```
myVal := []string{"val1", "val2"} 
myVal = append(myVal, "val3") //Appending a value to a sclice, actually recreates the slice

for i, myVal := range myVal { 
    fmt.Println(i, myVal)
}
```

Arrays and slices support range syntax as in Python: myArr[1:3] etc.

Internally slice is also saved as an array object only.
The slice variable points to an object in memory that has details about the slice like its length, and a pointer to the actual array object in the memory.

That's the reason why slice is pass by reference as opposed to array's pass by value. Basically the main data structure of slice having metadata about slice and the pointer to array in memory is copied over and passed over, but as you see, the underlying array is still the same. 

If you look at Reader interface: https://pkg.go.dev/io#Reader, it takes a byte slice as input, and in response it ony returns no of bytes read, and the error object. But where is the data read? The data is put into the the byte slice being passed as argument, and since it's pass by reference, the calling code gets the value in the object it had passed to the reader call.

```
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

### Maps
Key-value pairs. All keys are of same data type, and all values are of same data types. eg:
```
colorCodes := map[int]string { //this tells that keys are of type int and values are of type string
    1: "red",
    2: "blue"
}
```
Maps behave as pass by reference.

### Multiple return types

Go functions can return multiple return values:
```
func getEmployee() {
    return 'Ramit', 1
}

name, id = getEmployee()
```

### Custom types with receiver functions
Go doesn't support OOPS per say. But you can simulate the behaviour using custom types.

Custom types extend one of the basic data types of string, integer, float, map, array, slice.

Function with a receiver is like a 'method', a function that belongs to an 'instance'. 

See eg. of deck.go on how to define custom types with their reference functions.

### structs
Group data together along with receiver functions. 

All data members in struct are given default values as per their data type.

You can have structs within a struct. See employee.go for example.

To update values in struct, pass them around using pointers to mimic pass by reference.

Note the pointer shortcut supported by Go when calling reference functions on struct objects in main.go:

```
jim.updateName("jimmy") //is treated equivalet to below 2 lines:

jimPointer := &jim
jimPointer.updateName("jimmy")
```

### Type conversions
Go supports type conversions with valid data types. 
See deck.go for example.
We need to convert a custom type deck into a byte array for the method that saves data to file.
So basically we need to do conversions like this: deck -> []string -> string -> []byte.
In deck.go file, we do some of the conversions in toString function, and rest in saveToFile function.

### Unit testing
Run **go test** to run test cases.
Test files are named *_test.go. 
There are no explicit asserts, but we throw errors in test cases if results not as per expectations.
If no error thrown, then 'go test' would show 'PASS'.

### Interfaces
Interfaces are contractors that basically help to manage types.
There is no explicit 'implements interface'.
Any type that is implementing all functions (with same name and arguments) as the interface function, gets promoted 
to be of type implementing that interface.

Take a look at inheritance.go, on how the square struct is implementing shape interface, but rectange interface doesn't.

golang provided packages make heavy use of interfaces.

### Channels and Go Routines
Both are used for connurrency (with blocking code) in Go.

Use 'Go' keyword to launch a seperate thread (go Routine) to execute a function, seperate from the main calling thread.
```
go funcCall(args)
```

Channels are used to share data across go routines and the mmain calling code. Channels are defined of specific type, and only that kind of data can be put into it.

Channels can also be listened to in the main calling code to wait for all the threads to complete, else main code would exit before the threads launched finish.

```
channel <- 5 //send value 5 into a channel
myNumber <- channel //wait for a value to be assigned to a variable
fmt.Println(<-channel) //wait for a value to be sent into the channel. When we get one, log it out immediately.
```

See the example of sitechecker.go, where we check a group of websites if they are reachable or not, as separate threads, without blocking each other using go routines and channels.

Golang proves a Sleep() function which can be used to pause a function.

### Function literals

Function literals in Golang are also known as anonymous functions or lambdas. They are functions that are defined without a name and can be used as values or arguments. Function literals can also access variables from the surrounding scope, creating a closure. A closure is a function that can share and modify variables that are defined outside of its own scope.

```
func main() {
    // Define a function literal that takes two integers and returns their sum
    add := func(x, y int) int {
        return x + y
    }
    // Call the function literal with arguments 3 and 4
    result := add(3, 4)
    fmt.Println(result) // prints 7
}
```