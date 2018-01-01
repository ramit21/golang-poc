# golang-poc

Installing Golang: https://golang.org/doc/install

Learn about Golang: https://golang.org/

Steps to run the poc: (on Windows)

1) Ensure that Golang is installed and environment variables are set as explained in the installation link above

2) Git clone this repo

3) Execute the command 'go build' to create the executable

4) Run the executable: 'golang-poc'

5) Run the executable passing runtime arguments: 'golang-poc -name=Your_name'

6) Run 'golang-poc -h' to see the help menu with all the run time arguments defined for the application

GO 1.5 onwards supports cross compilation (https://dave.cheney.net/2015/03/03/cross-compilation-just-got-a-whole-lot-better-in-go-1-5)
To run the poc on OSX platform, follow the above steps with following exceptions:
In Step 3, generate the shell script using the command: 'env GOOS=darwin GOARCH=386 go build'

Now execute the script as './golang-poc -name=Your_name'
  
