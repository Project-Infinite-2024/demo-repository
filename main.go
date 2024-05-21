// main.go
package main

import "fmt"

// Add sums two integers and returns the result.
func Add(a, b int) int {
    return a + b
}

func main() {
    fmt.Println("2 + 3 =", Add(2, 3))
}
