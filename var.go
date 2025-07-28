package main

import (
	"fmt"
)

//
// -------------------------
// 📦 Package-Level Variables
// -------------------------

// Global variable with explicit type
var globalCount int = 100

// Inferred type
var appName = "Go Variable Demo"
var version = 1.0

// Grouped declaration
var (
	serverName string = "localhost"
	port       int    = 8080
)


// Constants (cannot be reassigned; immutable)
const maxError = 6
const (
	maxRetries=5
	PI=3.14
)



func main() {
	//
	// ----------------------------------------
	// 1. Explicit Declaration with Type
	// ----------------------------------------
	var username string = "gilfoyle"
	var name string = "Gopher"
	fmt.Println("Username:", username)
	fmt.Println("Name:", name)

	//
	// ----------------------------------------
	// 2. Implicit Type Inference using :=
	// (Only allowed inside functions)
	// ----------------------------------------
	retries := 3
	age := 10	//automatically infers int
	fmt.Println("Retries:", retries)
	fmt.Println("Age:", age)

	//
	// ----------------------------------------
	// 3. Multiple Variable Declarations
	// ----------------------------------------
	var firstName, lastName string = "Dinesh", "Chugtai"
	var x, y, z int = 1, 2, 3
	fmt.Println("Full name:", firstName, lastName)
	fmt.Println("x:", x, "y:", y, "z:", z)

	//
	// ----------------------------------------
	// 4. Zero Values (Uninitialized Vars, default values)
	// ----------------------------------------
	var userAge int        // defaults to 0
	var loggedIn bool      // defaults to false
	var isReady bool       // also defaults to false
	userAge=10
	fmt.Println("userAge:", userAge, "loggedIn:", loggedIn, "isReady:", isReady)

	//
	// ----------------------------------------
	// 5. Reassigning and Redeclaring Variables
	// ----------------------------------------
	username = "richard"    // reassignment
	x, newVar := 42, "new"  // short redeclaration within function scope
	fmt.Println("Updated Username:", username)
	fmt.Println("Updated x:", x, "newVar:", newVar)

	//
	// ----------------------------------------
	// 6. Block Declaration (Scoped Grouping)
	// ----------------------------------------
	var (
		server   = "localhost"
		localPort = 8080
		protocol = "http"
	)
	fmt.Printf("Server: %s://%s:%d\n", protocol, server, localPort)

	//
	// ----------------------------------------
	// 7. Accessing Global and Constant Values
	// ----------------------------------------
	fmt.Println("App:", appName, "Version:", version)
	fmt.Println("Global Count:", globalCount)	//accessing global variables
	fmt.Println("ServerName:", serverName, "Port:", port)	//accessing grouped vars
	fmt.Println("Max Retries Allowed:", maxRetries)
	fmt.Println("Constant PI:", PI)
}


