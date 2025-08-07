28.07.2025
# Go Init
install go locally
then start via:
```sh
go mod init go-playground
```

to run a file
```sh
go run [file_name]
go run main.go  # to direcrtly run the file
```

to first compile and then run binary file:
```sh
go build main.go
./main
```


# Values and Variables
various value type like `strings`, `integers`, `floats`, `booleans` etc

variables are explicitly declared and used by compiler to type-check function calls
```go
var a="initial"
var b,c int=1,2
var e                       //zero-valued variable, value is 0
fmt.Println(e)              //0
```

`zero-valued valriable`: Variable declared without corresponding initialization

`var.go`


# Constants
`const` declares a constant value.
```go
const s string = "constant"
```

	
A numeric constant has no type until it’s given one, such as by an explicit conversion.

# For
most basic type with a single condition
can have multiple condition also
iteration over range
```go
i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Another way of accomplishing the basic "do this
	// N times" iteration is `range` over an integer.
	for i := range 3 {
		fmt.Println("range", i)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

```

# If/Else
can skip else completely it want to, works without parantheses

yeah doesn't have any ternary operator
```go
package main

import "fmt"

func main() {

    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
```

# Switch Statement
switch is for like multiple if-else but in a beautiful way
```go
switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    default:
        fmt.Println("no num")
    }
```

You can use commas to separate multiple expressions in the same case statement.
```go
switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
```

[continue here](https://gobyexample.com/arrays)



--------------------------------------
[docs](https://go.dev/doc/tutorial/getting-started)
[tour](https://go.dev/tour/basics/1)
[web-dev in go](https://gowebexamples.com/)

[officially to go thru like book](https://gobyexample.com/)

[gobook](https://github.com/adonovan/gopl.io)

[cohort-docs](https://petal-estimate-4e9.notion.site/Golang-cohort-1257dfd1073580238258fe25973c319b)

