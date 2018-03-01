package main
import "fmt"

func fib_next(a, b int) int{
    return a+b
}

func main() {
    last := []int{1, 2}
    even := 2
    next := 0
    
    for i := 1; last[i] <= 4000000; i++{
        next = fib_next(last[i-1], last[i])
        last = append(last, next)
        if next%2==0{
            even += next
        }

    }
    fmt.Println(even)
}