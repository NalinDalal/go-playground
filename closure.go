package main
import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

   

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
		 nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())



    posSum := adder()
    fmt.Println(posSum(10)) // 10
    fmt.Println(posSum(5))  // 15
    fmt.Println(posSum(2))  // 17
}

