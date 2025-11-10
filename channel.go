package main

import "fmt"

func main() {

    messages := make(chan string)	//channel initialized

    go func() { messages <- "ping" }()	//value 'ping' sent into channel

    msg := <-messages
    fmt.Println(msg)
}
