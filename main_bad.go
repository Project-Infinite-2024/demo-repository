package main

import "fmt"

// Function to be tested
func Add(a, b int) int {
    return a + b
}

func main() {
    fmt.Println("2 + 3 =", Add(2, 3))
}
