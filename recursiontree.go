package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n // Base case: F(0) = 0, F(1) = 1
    }
    // Recursive step: F(n) = F(n-1) + F(n-2)
    return fibonacci(n-1) + fibonacci(n-2) 
}

func main() {
    n := 5
    fmt.Println("Fibonacci Numbers:")
    for i := 0; i < n; i++ {
        fmt.Println(fibonacci(i)) // Output: 0, 1, 1, 2, 3
    }
}
