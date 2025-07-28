package main

import (
	"fmt"
	"time"
"math/rand"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
}

