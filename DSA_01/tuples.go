package main 

import "fmt"

func main() {
    
    var square int
    var cube int
    square, cube = powerSeries(69)
    fmt.Printf("Square: %v, Cube: %v", square, cube)
}

func powerSeries(n int) (int,int) {
    return n*n, n*n*n
}


