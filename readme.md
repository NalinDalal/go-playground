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

---

# Generics

Starting with **Go 1.18**, Go introduced support for **generics**, also known as **type parameters**.
Generics allow functions and types to operate on values of different types **without duplicating code**.

### Example: Generic Function

```go
package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}

func main() {
    var s = []string{"foo", "bar", "zoo"}

    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
    _ = SlicesIndex[[]string, string](s, "zoo") // explicit type args
}
```

### Explanation

- **Type parameters** appear inside square brackets `[S ~[]E, E comparable]`.
- `S ~[]E` means: `S` is a slice type whose element type is `E`.
- `E comparable` constrains `E` to types that can be compared using `==` and `!=`.
- The function iterates through the slice and returns the **index** of the first matching element.
- Returns `-1` if not found.
- The **compiler can infer** type parameters in most cases, so explicit type arguments are usually not needed.

  **Note:**
  A similar function already exists in Go’s standard library as `slices.Index`.

### Example: Generic Type

```go
type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}
```

This defines a **singly-linked list** whose elements can be of **any type** `T`.

### Methods on Generic Types

We can define methods on generic types the same way as on regular types,
but we must **include the type parameter** when defining the receiver.

```go
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}
```

### Using the Generic List

```go
func main() {
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())
}
```

### Key Points

- **Generics** enable writing reusable, type-safe code without sacrificing performance.
- **Type parameters** are defined in square brackets `[T any]`.
- **Constraints** (`comparable`, `~[]E`, etc.) define what operations are allowed on the type.
- **Type inference** means you often don’t need to explicitly specify type parameters.
- You can create **generic functions**, **generic types**, and **methods** on them.
- The type name with its parameters (e.g. `List[T]`) must always include the `[T]` when defining methods.

---

# Range over Iterators

previous example: `AllElements` method that returned a slice of all elements in the list.
we can do better

Starting with **Go 1.23**, Go added **iterators**, allowing us to **range over almost any sequence** including linked lists, streams, or even infinite sequences without converting to a slice first.

Iterators in Go are **functions with a specific signature**, which take a `yield` callback function. `yield` is called for each element in the sequence, and can signal **early termination** by returning `false`.

```go
//Example: Iterator for a Linked List
type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

// Returns an iterator
func (lst *List[T]) All() iter.Seq[T] {
    return func(yield func(T) bool) {
        for e := lst.head; e != nil; e = e.next {
            if !yield(e.val) {
                return
            }
        }
    }
}
```

## How It Works

- `iter.Seq[T]` is the **iterator type**, a function that accepts `yield func(T) bool`.
- Each element is sent to `yield`.
- If `yield` returns `false`, iteration stops early.
- No intermediate slice is needed — iteration is **lazy** and memory-efficient.

## Using the Iterator

```go
func main() {
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)

    // Iterate using range
    for e := range lst.All() {
        fmt.Println(e)
    }

    // Collect all values into a slice
    all := slices.Collect(lst.All())
    fmt.Println("all:", all)
}
```

## Infinite Iterators Example (Fibonacci)

```go
func genFib() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 1, 1
        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}

func main() {
    for n := range genFib() {
        if n > 50 { // early exit
            break
        }
        fmt.Println(n)
    }
}
```

**Key Points**

- Iterators allow **lazy iteration** over any sequence, including infinite ones.
- The `yield` function controls both output and early termination.
- Packages like `slices` provide helper functions like `Collect` to gather iterator values into slices.
- Iterators remove the need for intermediate allocations like `AllElements()` slices.

---

# Errors

Error handling is accounting for scenarios when the code fails

By convention, errors are the last return value and have type error, a built-in interface.

example:

```go
func f(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }
    return arg + 3, nil
}
```

`errors.New` constructs a basic error value with the given error message.
example:

```go
return errors.New("can't work with 42")
```

A `nil` value in the error position indicates that there was no error.

```go
return arg + 3, nil
```

`sentinel` error: predeclared variable that is used to signify a specific error condition.

```go
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")
```

Note: you can wrap errors with error to add context, it's cause Wrapped errors create a logical chain (A wraps B, which wraps C, etc.) .

```go
func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {

        return fmt.Errorf("making tea: %w", ErrPower)
    }
    return nil
}
```

idomatically you should use errors in if-else blocks

```go
for _, i := range []int{7, 42} {

        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }
```

`errors.Is` checks that given error matches a specific error value.
useful with wrapped or nested errors,

```go
if errors.Is(err, ErrOutOfTea) {
    fmt.Println("We should buy new tea!")
} else if errors.Is(err, ErrPower) {
    fmt.Println("Now it is dark.")
} else {
    fmt.Printf("unknown error: %s\n", err)
}
```

# Errors : Custom

define custom error types by implementing the `Error()` method on them
conventions is to have suffix `Error` in nomenclature, ex: `argError`

```go
//declaration
type argError struct {
    arg     int
    message string
}

// implementing interface
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}


func f(arg int) (int, error) {
    if arg == 42 {
        //Return our custom error.

        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}
```

`errors.As` is a more advanced version of `errors.Is`
checks that a given error (or any error in its chain) matches a specific error type and converts to a value of that type, returning true.
If there’s no match, it returns false.

---

# Go Routines

lightweight thread of execution.

say we have a function `f(s)` we call it normally in main function like

```go
func main() {

    f("direct")
    //runs synchornously
}
```

but to invoke in `go-routines` we do like:

```go
func main(){
    go f("goroutine")
    //execute concurrently with the calling one.
}
```

can do same for anonymous functions

```go
go func(msg string) {
        fmt.Println(msg)
    }("going")
```

both are calling asynchonosuly so we wait for them to finish up

```go
time.Sleep(time.Second)
    fmt.Println("done")
```

---

# Channels

pipes that connect concurrent goroutine
it's like channel has 2 endpoint, both have a goroutine at them
1 will send data and another will receive it

initialize channel:
`make(chan val-type)` ; typed by values they convey
Send a value into a channel using the `channel <-` syntax.

The `<-channel` syntax $$receives$$ a value from the channel.

```go

    messages := make(chan string)   //channel initialized

    go func() { messages <- "ping" }()  //value 'ping' sent into channel

    msg := <-messages
    fmt.Println(msg)
```

By default sends and receives block until both the sender and receiver are ready. so no need to worry about synchronisation.

## Channel Buffering

by default channels are `unbuffered` i.e. they will only receive data

Buffered channels accept a limited number of values without a corresponding receiver for those values.

```go
make(chan string, 2)    //channel made for strings Buffering upto 2 values
//can send data without concurrent receiver
messages <- "buffered"
messages <- "channel"
```

## Channel Synchronisation

channels can be used to synchronis the executions across goroutines.
take an example where you can block `receive` to wait for multiple goroutines to finish.

```go
package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done
}
```

`done:=make(chan bool,1)` a worker is made and then you gave it channel to work on.

`<-done` basically does what is tells the program to wait until worker is done
if we remove it then it will exit directly, without waiting for worker to finish.

## Channel Direction

well you can specify if channel will send or receive
increases type-safety

we use spcl words `ping` to receive, and `pong` to send/transmit data
need to give out this in function arguments

```go
func ping(pings chan<- string, msg string) {
    //single channel to send only
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    //accepts 1 channel to receive, 1 to send
    msg := <-pings
    pongs <- msg
}
```

if you try to send data when only receive is argued then it will create compile time errors.

## Select

`select` lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful feature of Go.
allows a goroutine to wait on multiple channel operations simultaneously.

```go
 // Two channels for concurrent communication
    c1 := make(chan string)
    c2 := make(chan string)

    // Simulate two concurrent operations
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // Wait for both results using select
    for range 2 {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
```

- Each goroutine sends a message after a delay (`1s` and `2s`) — simulating **blocking RPC calls** or **network operations**.
- The `select` statement waits on **multiple channel operations** and executes the **first one that becomes ready**.
- When multiple cases are ready, one is chosen **at random**.

well you can time the process:

```sh
time go run select.go
```

Output:

```
received one
received two
```

Total time ≈ **2 seconds**, since both goroutines sleep and send concurrently.

### Key Points

- `select` listens on **multiple channel operations** (`send` or `receive`).
- It **blocks** until one of the cases can proceed.
- When a case executes, the others are **ignored** for that iteration.
- Great for handling **multiple concurrent results**, **timeouts**, or **fan-in** patterns.

### Use Cases

- Waiting for the **first available result** among multiple goroutines.
- Implementing **timeouts** or **cancellation** logic.
- **Merging multiple input channels** into a single consumer.

> `select` is like a `switch`, but for **channels** — letting Go routines communicate efficiently and handle whichever channel becomes ready first.

---

# Timeout

helps connect to external resources or to bound execution time for long-running tasks.
They help ensure your program doesn’t wait indefinitely for a response.

we implement them using `channels` and `select`

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Simulate an external call that takes 2 seconds to complete
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    // Wait for result or timeout after 1 second
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    // Another simulated call
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()

    // This time, allow enough time for the operation to finish
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}
```

- **`time.After(d)`** returns a channel that sends a value after `d` duration.
- In each `select`, we wait for either:
  - the result from the worker goroutine, or
  - the timeout signal from `time.After`.

- Whichever channel becomes **ready first** determines which case executes.

### Output

```
timeout 1
result 2
```

**Explanation of output:**

- The first operation takes 2 seconds but we only wait 1 second → timeout triggered.
- The second operation also takes 2 seconds, but the timeout is 3 seconds → result received successfully.

### Why Buffered Channels?

Each worker uses a **buffered channel** (`make(chan string, 1)`):

- Prevents a **goroutine leak** in case the result is never read.
- Without buffering, the send (`c1 <- "result 1"`) could block forever if no receiver is ready.

### Key Points

- `time.After(duration)` provides a **timeout signal** via channel.
- `select` ensures your code reacts to **whichever event happens first**.
- Use buffered channels to **avoid blocking** when sending results from background goroutines.
- Commonly used for **network calls**, **API requests**, and **resource polling**.

> Combine `select` + `time.After` to handle operations that might take too long — safely and cleanly.

---

[continue](https://gobyexample.com/non-blocking-channel-operations)

[docs](https://go.dev/doc/tutorial/getting-started)
[tour](https://go.dev/tour/basics/1)
[tcp-to-http](https://www.youtube.com/watch?v=FknTw9bJsXM&list=WL&index=1&t=18s)
[web-dev in go](https://gowebexamples.com/)

[cohort-docs](https://petal-estimate-4e9.notion.site/Golang-cohort-1257dfd1073580238258fe25973c319b)
