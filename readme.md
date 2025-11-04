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

# Array

an array is a numbered sequence of elements of a specific length.

```go
    var a [5]int    //declaration
    b := [5]int{1, 2, 3, 4, 5}
    b = [...]int{1, 2, 3, 4, 5}

    //2d array
    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }

```

then you have some of functions such as `set`, `get`
We can set a value at an index using the `array[index] = value` syntax, and get a value with `array[index]`.

```go
a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))
```

now if you specify the index with `:` so it will put elements at that index, rest will
be initialised to `0`

```go
b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)
```

array are suppossed to be of single dimension, but you can support multi-dimension too

```go
twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
fmt.Println("2d: ", twoD)
```

# Slices

slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0.
To create a slice with non-zero length, use the builtin make.

```go
var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
```

we make a slice of strings of length 3 (initially zero-valued). By default a new slice’s capacity is equal to its length; if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to make.

```go
s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
```

there are more of methods which make slices richer than arrays

- append method: returns a slice containing one or more new values

```go
s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)
```

- copy : copy into c from s.

```go
c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)
```

- slices can also be used as similar to python syntax

```go
l := s[2:5]
    fmt.Println("sl1:", l)
```

- slices up to (but excluding) s[5].

```go
l = s[:5]
    fmt.Println("sl2:", l)
```

- slices up from (and including) s[2].

```go
l = s[2:]
    fmt.Println("sl3:", l)
```

- declare and initialize a variable for slice in a single line

```go
t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
```

- other utility functions:

```go
t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }
```

- working with multi-dimension data

```go
twoD := make([][]int, 3)
    for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
```

# Maps

built-in associative data types

use the builtin make: `make(map[key-type]val-type)`.

Set key/value pairs using typical `name[key] = val` syntax.

```go
m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13
```

to print all of values:

```go
fmt.Println("map:", m)
```

get a value for a key with `name[key]`

```go
    v3 := m["k3"]
    fmt.Println("v3:", v3)
```

`len(m)`- length of map

builtin delete removes key/value pairs from a map.

```go
delete(m, "k2")
    fmt.Println("map:", m)
```

remove all key/value pairs from a map, use the clear builtin.

```go
clear(m)
    fmt.Println("map:", m)
```

declare and initialize a new map in the same line with this syntax.

```go
n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
```

some utility function:

```go
n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
```

# Function

function are something which takes some input and produce some output
ex: a function that takes two ints and returns their sum as an int.

```go
func plus(a int, b int) int {
    return a + b
}

func main(){
    res := plus(1, 2)
    fmt.Println("1+2 =", res)
}
```

When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.

```go
func plusPlus(a, b, c int) int {
    return a + b + c
}
```

Call a function just as you’d expect, with name(args).

## Multiple Return Value

built-in support for multiple return values
useful in idomatic go where you return both result and error

```go
func vals() (int, int) {        //(int, int) in this function signature shows that the function returns 2 ints.
    return 3, 7
}
```

If you only want a subset of the returned values, use the blank identifier \_.

```go
_, c := vals()
    fmt.Println(c)
```

# Variadic Functions

can be called with any number of trailing arguments.

Variadic functions can be called in the usual way with individual arguments.

```go
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
//Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
```

# Closures

A closure is a function value that captures variables from its surrounding scope. Even after the surrounding function has returned, the closure can still access and modify those captured variables.

Here functions are first-class citizens, so you can define and return functions from other functions. Closures make use of this.

1. A closure is created when an inner function **references variables defined outside its body**.

2. The captured variables are **shared** across all closures created in that scope.

3. Closures are often used to **maintain state** between function calls without using global variables.

```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```

# Recursion

function that calls itself

```go
var fib func(n int) int
    fib = func(n int) int {
        if n < 2 {
            return n
        }
return fib(n-1) + fib(n-2)
    }
```

# Range over Built-in Types

range iterates over elements in a variety of built-in data structures.

```go
//over array
for _, num := range nums {  //index ignored
        sum += num
    }

for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
```

```go
//over maps
kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

//over keys of map
for k := range kvs {
        fmt.Println("key:", k)
    }
```

# Pointers

allows to pass references to values and records within the program

pass by value

```go
func zeroval(ival int) {    //parameter is int
    ival = 0
}
```

pass by pointer

```go
func zeroptr(iptr *int) {
    *iptr = 0
}
```

`zeroptr` in contrast has an `*int` parameter, meaning that it takes an int pointer.

`*iptr` _dereferences_ the pointer from its memory address to the current value at that address.

`&i` gives the memory address of i
i.e. a pointer to i

```go
zeroptr(&i)
fmt.Println("zeroptr:", i)
```

pointers can be printed too

```go
fmt.Println("pointer:", &i)
```

---

# Strings and Runes

Strings in Go are **read-only slices of bytes**.
They are containers of UTF-8 encoded text, but indexing a string gives you **bytes**, not “characters.”

A **rune** is Go’s term for a **Unicode code point** (type `int32`).

### String basics

```go
const s = "สวัสดี" // "hello" in Thai
```

`len(s)` → gives the number of **bytes**, not runes.

```go
fmt.Println("Len:", len(s))  // 18
```

Indexing a string gives the raw **byte values**:

```go
for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i])
}
// e0 b8 aa e0 b8 a7 ...
```

### Counting runes

Use the `utf8` package to count runes (decoded code points):

```go
import "unicode/utf8"

fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 6
```

### Iterating over runes

`range` on a string decodes runes automatically:

```go
for idx, r := range s {
    fmt.Printf("%#U starts at %d\n", r, idx)
}
```

Output:

```
U+0E2A 'ส' starts at 0
U+0E27 'ว' starts at 3
...
```

Equivalent explicit decoding:

```go
for i, w := 0, 0; i < len(s); i += w {
    r, w := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%#U starts at %d\n", r, i)
}
```

### Rune literals

Runes use **single quotes**:

```go
func examineRune(r rune) {
    if r == 'ส' {
        fmt.Println("found so sua")
    } else if r == 't' {
        fmt.Println("found tee")
    }
}
```

### Key points

- A **string** is a sequence of bytes.
- String **literals** are UTF-8 by default.
- A **rune** = Unicode code point (`int32`).
- `len(string)` counts **bytes**, not runes.
- Use `utf8.RuneCountInString` or `range` for runes.
- Indexing gives **bytes**, not runes.

---

# Structs

typed collections of fields. They’re useful for grouping data together to form records.

ex:

person struct type has name and age fields.

```go
type person struct {
    name string
    age  int
}
```

`newPerson` constructs a new person struct with the given name.

```go
func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}
```

Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it.

- can name fields while initialising
- Omitted fields will be zero-valued.
- An `&` prefix yields a pointer to the struct.

idiomatic to encapsulate new struct creation in constructor functions

Access struct fields with a dot.

```go
s := person{name: "Sean", age: 50}
    fmt.Println(s.name)
```

use dots with struct pointers - the pointers are automatically dereferenced.

```go
sp := &s
    fmt.Println(sp.age)
```

they mutable

when using for single value, no need to give it name; it can have anonymous struct type.

```go
dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
fmt.Println(dog)
```

---

# Methods

go supports method defined on struct types

```go

func (r *rect) area() int {
    return r.width * r.height
}
```

This area method has a receiver type of \*rect.

Methods can be defined for either pointer or value receiver types.

```go
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}
```

Go automatically handles conversion between values and pointers for method calls.
You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.

---

# Interfaces

Interfaces are named collections of method signatures.

ex:

```go
type geometry interface {
    area() float64
    perim() float64
}
type rect struct {
    width, height float64
}
//implement this interfaces on rectangle type
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
```

To implement an interface in Go, we just need to implement all the methods in the interface.

If a variable has an interface type, then we can call methods that are in the named interface.

```go
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
```

# Enumerated Types (Enums)

Go doesn’t have a built-in `enum` type, but enums can be implemented using constants and the `iota` keyword.

### Example

```go
package main

import "fmt"

// Define an enum-like type with underlying int
type ServerState int

// Define possible values using iota
const (
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)
```

`iota` generates successive constant values automatically (0, 1, 2, ...).

### Adding String Representation

To make enum values printable, implement the `fmt.Stringer` interface:

```go
var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
    return stateName[ss]
}
```

This allows direct printing of `ServerState` values, e.g. `fmt.Println(StateIdle)` → `idle`.

If there are many enum values, this can be automated with the `stringer` tool using `go:generate`.

### Using the Enum

```go
func main() {
    ns := transition(StateIdle)
    fmt.Println(ns)
    ns2 := transition(ns)
    fmt.Println(ns2)
}
```

### State Transition Example

```go
func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        // logic for transitioning back to idle
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}
```

### Key Points

- Enums in Go are implemented using **typed constants**.
- `iota` automatically increments values within a constant block.
- Implementing `String()` improves readability and debuggability.
- Using a distinct type (e.g., `ServerState int`) enforces **compile-time type safety** — you can’t pass a plain `int` where a `ServerState` is expected.

---

# Struct Embedding

embedding of structs and interfaces to express a more seamless composition of types.

A container embeds a base. An embedding looks like a field without a name.

```go
type container struct {
    base
    str string
}
```

When creating structs with literals, we have to initialize the embedding explicitly;

```go
co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }
```

can access fields direcrtly like: `co.base.num` or `co.str`

Since container embeds base, the methods of base also become methods of a container.

```go
fmt.Println("describe:", co.describe())
```

[continue here](https://gobyexample.com/generics)

---

[docs](https://go.dev/doc/tutorial/getting-started)
[tour](https://go.dev/tour/basics/1)
[tcp-to-http](https://www.youtube.com/watch?v=FknTw9bJsXM&list=WL&index=1&t=18s)
[web-dev in go](https://gowebexamples.com/)

[officially to go thru like book](https://gobyexample.com/)

[gobook](https://github.com/adonovan/gopl.io)

[cohort-docs](https://petal-estimate-4e9.notion.site/Golang-cohort-1257dfd1073580238258fe25973c319b)
