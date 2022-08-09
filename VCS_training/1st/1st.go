//open terminal
// go run is to compiles and execute one or two files
//go build is just to compile
// go fmt (format) format all the code in each line in the current directory
//go install is to compile and install a package
// go test is to download the raw source code of someone else's package
//go test is to runs any test associated with the current project

package main //--> main.exe --> executed package
//others name is dependency type
// package = project or workspace
//need to declare package for any code in the package
//main is the name of the package that we define its a executable or reusable

import "fmt"

// give my package main access to all of the code and functionality that is contained inside of this package"fmt"
// fmt = standard library package ==format

func main() {
	// func is to tells go we are about to declare a function
	// main is to set the name of the function
	//() is the list of arguments to pass the function
	fmt.Println("Hi there")
}

// 1st declare package
//2nd import other packeages that we need
// declare functions, tell Go to do things
