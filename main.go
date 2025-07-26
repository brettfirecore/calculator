package main

import "fmt"

func main() {
    // Declare variables with types
    var title string
    var price float64
    var inStock bool
    var copies int

    // Assign values to those variables
    title = "For the Love of Go"    // variable name on the left hand side of =
    price = 39.99                   // & value on the right hand side
    inStock = true                  // bool is true or false
    copies = 12                     // int or float64 whole or decimal number

    // Print the values
    fmt.Println("Book title:", title)
    fmt.Println("Price:", price)
    fmt.Println("In stock:", inStock)
    fmt.Println("Number of copies:", copies)
}
