package main

import "fmt"

func main() {
    var title string
    var copies int
    var onSpecialOffer bool
    var discountPercentage float64  // declare this variable
    var edition int                 // new variable for edition number

    title = "For the Love of Go"
    copies = 99
    onSpecialOffer = true           // Book is on special offer
    discountPercentage = 10.0       // 10% discount
    edition = 2                     // assign suitable edition number

    fmt.Println(title)
    fmt.Println(copies)
    fmt.Println("On special offer:", onSpecialOffer)
    fmt.Println("Discount percentage:", discountPercentage)
    fmt.Println("Edition:", edition)  // print edition number
}
