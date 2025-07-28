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
![continue here](https://gobyexample.com/constants)



--------------------------------------
![docs](https://go.dev/doc/tutorial/getting-started)
![tour](https://go.dev/tour/basics/1)
![web-dev in go](https://gowebexamples.com/)

![officially to go thru like book](https://gobyexample.com/)

![gobook](https://github.com/adonovan/gopl.io)

