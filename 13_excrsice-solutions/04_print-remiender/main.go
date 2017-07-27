package main

import "fmt"

func main() {
  var numerator, denominator int
  fmt.Print("Enter the numerator: ")
  fmt.Scan(&numerator)
  fmt.Print("Enter the denominator: ")
  fmt.Scan(&denominator)
  fmt.Println(numerator % denominator)
}