# golang-poc

Installing Golang: https://golang.org/doc/install

Learn about Golang: https://golang.org/

Steps to run the poc: (on Windows)

1) Ensure that Golang is installed and environment variables are set as explained in the installation link above

2) Git clone this repo

3) Execute the command 'go build' to create the executable

4) Run the executable: 'm.exe'

5) Run the executable passing runtime arguments: 'golang-poc -name=Your_name'

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

### Multiple return types

Go functions can return multiple return values:
```
func getEmployee()
```

### Custom types with receiver functions
Go doesn't support OOPS per say. But you can simulate the behaviour using custom types.

Custom types extend one of the basic data types of string, integer, float, map, array, slice.

Function with a receiver is like a 'method', a function that belongs to an 'instance'. 

See eg. of deck.go on how to define custom types with their reference functions.